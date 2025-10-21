//go:build examples

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/custom_objects"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	baseURL := os.Getenv("ACTIVE_URL")
	token := os.Getenv("ACTIVE_TOKEN")

	c, err := client.NewCoreClient(baseURL, token)
	if err != nil {
		log.Fatalf("failed to create core client: %v", err)
	}

	svc := custom_objects.NewRealService(c)

	log.Printf("Using ACTIVE_URL=%s", baseURL)
	if token == "" {
		log.Printf("ACTIVE_TOKEN not set")
	} else {
		log.Printf("ACTIVE_TOKEN appears set (redacted)")
	}

	if baseURL == "" {
		log.Printf("ACTIVE_URL not set; example will still run but will likely return not implemented errors")
	}

	// Safe mode: by default the example will NOT perform create/update/delete.
	// To allow destructive actions set CUSTOM_OBJECTS_SAFE=false in the environment.
	safeEnv := os.Getenv("CUSTOM_OBJECTS_SAFE")
	safeMode := true
	if strings.ToLower(safeEnv) == "false" || safeEnv == "0" {
		safeMode = false
	}
	if safeMode {
		log.Printf("Running in SAFE mode: create/update/delete will be skipped. Set CUSTOM_OBJECTS_SAFE=false to allow destructive actions.")
	}

	// Optional relationship configuration: set CUSTOM_OBJECT_REL_NS and CUSTOM_OBJECT_REL_ID
	// e.g. CUSTOM_OBJECT_REL_NS=contacts CUSTOM_OBJECT_REL_ID=22
	relNS := os.Getenv("CUSTOM_OBJECT_REL_NS")
	relID := os.Getenv("CUSTOM_OBJECT_REL_ID")

	// 1) List object types (include fields)
	listResp, apiResp, err := svc.ListObjectTypes(context.Background(), map[string]string{"showFields": "all"})
	if apiResp != nil {
		log.Printf("ListObjectTypes HTTP status: %d", apiResp.StatusCode)
		if len(apiResp.Body) > 0 {
			log.Printf("ListObjectTypes raw body:\n%s", string(apiResp.Body))
		}
	}
	if err != nil {
		log.Printf("ListObjectTypes error: %v", err)
	} else {
		log.Printf("ListObjectTypes: found %d schemas", len(listResp.Schemas))
		if apiResp != nil && len(apiResp.Body) > 0 {
			var raw map[string]interface{}
			_ = json.Unmarshal(apiResp.Body, &raw)
			// Compare shapes roughly
			var typed custom_objects.ListSchemasResponse
			_ = json.Unmarshal(apiResp.Body, &typed)
			if reflect.DeepEqual(typed, *listResp) {
				log.Printf("raw JSON and typed response match for ListObjectTypes")
			} else {
				log.Printf("raw != typed for ListObjectTypes: raw keys=%v, typed schemas=%d", reflect.ValueOf(raw).MapKeys(), len(listResp.Schemas))
			}
		}
	}

	// 2) Optionally, create a temporary object type (cleanup later). Only run when safeMode == false
	var createdSchema *custom_objects.Schema
	// track IDs for deferred cleanup when safeMode == false
	var createdSchemaID string
	var createdRecordIDs []string
	// track which created IDs were already deleted during the main run so
	// deferred cleanup won't attempt to delete them again (avoids 404s).
	createdRecordDeleted := make(map[string]bool)
	var createdSchemaDeleted bool
	if !safeMode {
		// Create a minimal schema with one text field
		schemaReq := &custom_objects.CreateObjectTypeRequest{
			Name:   "ExampleType",
			Slug:   "example-type-demo",
			Labels: &custom_objects.LabelPair{Singular: "Example Type", Plural: "Example Types"},
			Fields: []custom_objects.SchemaField{
				{ID: "name", Labels: custom_objects.LabelPair{Singular: "Name", Plural: "Names"}, Type: "text", Required: false},
			},
		}
		createdSchema, apiResp, err = svc.CreateObjectType(context.Background(), schemaReq)
		if apiResp != nil {
			log.Printf("CreateObjectType HTTP status: %d", apiResp.StatusCode)
			if len(apiResp.Body) > 0 {
				log.Printf("CreateObjectType raw body:\n%s", string(apiResp.Body))
			}
		}
		if err != nil {
			log.Printf("CreateObjectType error: %v", err)
		} else {
			log.Printf("Created object type id=%s slug=%s", createdSchema.ID, createdSchema.Slug)
			createdSchemaID = createdSchema.ID
		}

		// Ensure we attempt cleanup at the end of the run even if an intermediate
		// step fails. This will attempt to delete created records and the
		// schema. Errors during cleanup are logged but do not halt the example.
		defer func() {
			if len(createdRecordIDs) > 0 {
				for _, rid := range createdRecordIDs {
					if createdRecordDeleted[rid] {
						// already deleted during main run
						continue
					}
					apiResp, err := svc.DeleteObjectRecord(context.Background(), createdSchemaID, rid)
					if apiResp != nil {
						log.Printf("Deferred DeleteObjectRecord HTTP status: %d", apiResp.StatusCode)
						if len(apiResp.Body) > 0 {
							log.Printf("Deferred DeleteObjectRecord raw body:\n%s", string(apiResp.Body))
						}
					}
					if err != nil {
						log.Printf("Deferred DeleteObjectRecord error for id=%s: %v", rid, err)
					} else {
						log.Printf("Deferred deleted record id=%s", rid)
					}
				}
			}
			if createdSchemaID != "" && !createdSchemaDeleted {
				apiResp, err := svc.DeleteObjectType(context.Background(), createdSchemaID)
				if apiResp != nil {
					log.Printf("Deferred DeleteObjectType HTTP status: %d", apiResp.StatusCode)
					if len(apiResp.Body) > 0 {
						log.Printf("Deferred DeleteObjectType raw body:\n%s", string(apiResp.Body))
					}
				}
				if err != nil {
					log.Printf("Deferred DeleteObjectType error for id=%s: %v", createdSchemaID, err)
				} else {
					log.Printf("Deferred deleted object type id=%s", createdSchemaID)
				}
			}
		}()
	} else {
		log.Printf("Skipping CreateObjectType because SAFE mode is enabled")
	}

	// 3) Get the type (if created)
	if createdSchema != nil && createdSchema.ID != "" {
		got, apiResp, err := svc.GetObjectType(context.Background(), createdSchema.ID)
		if apiResp != nil {
			log.Printf("GetObjectType HTTP status: %d", apiResp.StatusCode)
			if len(apiResp.Body) > 0 {
				log.Printf("GetObjectType raw body:\n%s", string(apiResp.Body))
			}
		}
		if err != nil {
			log.Printf("GetObjectType error: %v", err)
		} else {
			log.Printf("GetObjectType returned id=%s slug=%s", got.ID, got.Slug)
		}

		// 4) Create two records for the type
		for i, val := range []string{"demo-1", "demo-2"} {
			req := &custom_objects.CreateRecordRequest{Fields: map[string]interface{}{"name": val}}
			if relNS != "" && relID != "" {
				// Relationships are typed as map[string][]interface{} in the model.
				// Convert single relID value into a single-element slice.
				req.Relationships = map[string][]interface{}{relNS: []interface{}{relID}}
			}
			recResp, apiResp, err := svc.CreateObjectRecord(context.Background(), createdSchema.ID, req)
			if apiResp != nil {
				log.Printf("CreateObjectRecord HTTP status: %d", apiResp.StatusCode)
				if len(apiResp.Body) > 0 {
					log.Printf("CreateObjectRecord raw body:\n%s", string(apiResp.Body))
				}
			}
			if err != nil {
				log.Printf("CreateObjectRecord error: %v", err)
				continue
			}
			createdRecordIDs = append(createdRecordIDs, recResp.Record.ID)
			log.Printf("Created record #%d id=%s", i+1, recResp.Record.ID)
		}

		// 5) List records for the schema to confirm (request a larger limit)
		listRecs, apiResp, err := svc.ListObjectRecords(context.Background(), createdSchema.ID, map[string]string{"limit": "100"})
		if apiResp != nil {
			log.Printf("ListObjectRecords HTTP status: %d", apiResp.StatusCode)
			if len(apiResp.Body) > 0 {
				log.Printf("ListObjectRecords raw body:\n%s", string(apiResp.Body))
			}
		}
		if err != nil {
			log.Printf("ListObjectRecords error: %v", err)
		} else {
			log.Printf("ListObjectRecords: found %d records", len(listRecs.Records))
		}

		// 6) Get each created record by id
		for _, recordID := range createdRecordIDs {
			gotRec, apiResp, err := svc.GetObjectRecord(context.Background(), createdSchema.ID, recordID)
			if apiResp != nil {
				log.Printf("GetObjectRecord HTTP status: %d", apiResp.StatusCode)
				if len(apiResp.Body) > 0 {
					log.Printf("GetObjectRecord raw body:\n%s", string(apiResp.Body))
				}
			}
			if err != nil {
				log.Printf("GetObjectRecord error: %v", err)
			} else {
				log.Printf("Got record id=%s fields=%v", gotRec.ID, gotRec.Fields)
			}

			// 7) Update the record
			// Include the record ID in the update payload as some API versions expect it
			updReq := &custom_objects.UpdateRecordRequest{ID: recordID, Fields: map[string]interface{}{"name": "demo-updated"}}
			upd, apiResp, err := svc.UpdateObjectRecord(context.Background(), createdSchema.ID, recordID, updReq)
			if apiResp != nil {
				log.Printf("UpdateObjectRecord HTTP status: %d", apiResp.StatusCode)
				if len(apiResp.Body) > 0 {
					log.Printf("UpdateObjectRecord raw body:\n%s", string(apiResp.Body))
				}
			}
			if err != nil {
				log.Printf("UpdateObjectRecord error: %v", err)
			} else {
				log.Printf("Updated record id=%s", upd.Record.ID)
			}

			// 8) Delete the record
			apiResp, err = svc.DeleteObjectRecord(context.Background(), createdSchema.ID, recordID)
			if apiResp != nil {
				log.Printf("DeleteObjectRecord HTTP status: %d", apiResp.StatusCode)
				if len(apiResp.Body) > 0 {
					log.Printf("DeleteObjectRecord raw body:\n%s", string(apiResp.Body))
				}
			}
			if err != nil {
				log.Printf("DeleteObjectRecord error: %v", err)
			} else {
				log.Printf("Deleted record id=%s", recordID)
				// mark as deleted so deferred cleanup will skip it
				createdRecordDeleted[recordID] = true
			}
		}

		// 8) Delete the created type
		apiResp, err = svc.DeleteObjectType(context.Background(), createdSchema.ID)
		if apiResp != nil {
			log.Printf("DeleteObjectType HTTP status: %d", apiResp.StatusCode)
			if len(apiResp.Body) > 0 {
				log.Printf("DeleteObjectType raw body:\n%s", string(apiResp.Body))
			}
		}
		if err != nil {
			log.Printf("DeleteObjectType error: %v", err)
		} else {
			log.Printf("Deleted object type id=%s", createdSchema.ID)
			// mark schema as deleted to avoid deferred second-delete
			createdSchemaDeleted = true
		}
	}

	fmt.Println("example finished")
}
