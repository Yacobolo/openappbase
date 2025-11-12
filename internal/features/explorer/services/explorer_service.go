// internal/features/explorer/services/explorer_service.go
package services

// type ExplorerState struct {
// 	ActiveSchema string `json:"active_schema"`
// 	ActiveTable  string `json:"active_table"`
// }
// type ExplorerService struct {
// 	q     *store.Queries
// 	kv    jetstream.KeyValue
// 	state *session.StateStore[ExplorerState]
// }

// func NewExplorerService(q *store.Queries, js jetstream.JetStream, store sessions.Store) (*ExplorerService, error) {

// 	// Create a specific KV bucket for the explorer's state
// 	kv, err := js.CreateOrUpdateKeyValue(context.Background(), jetstream.KeyValueConfig{
// 		Bucket:      "explorer_state",
// 		Description: "State for the Unity Explorer",
// 		Compression: true,
// 		TTL:         24 * time.Hour, // Remember state for a day
// 	})
// 	if err != nil {
// 		return nil, fmt.Errorf("error creating key value for explorer: %w", err)
// 	}

// 	stateStore := session.NewStateStore[ExplorerState](kv, store)

// 	return &ExplorerService{
// 		q:     q,
// 		kv:    kv,
// 		state: stateStore,
// 	}, nil
// }

// func (s *ExplorerService) GetSessionID(r *http.Request, w http.ResponseWriter) (string, error) {
// 	return s.state.GetSessionID(r, w)
// }

// func (s *ExplorerService) GetState(ctx context.Context, sessionID string) (*ExplorerState, error) {
// 	state, err := s.state.Get(ctx, sessionID)
// 	// Add feature-specific default logic if needed
// 	if state.ActiveSchema == "" {
// 		state.ActiveSchema = "public"
// 	}
// 	return state, err
// }

// func (s *ExplorerService) SaveState(ctx context.Context, sessionID string, state *ExplorerState) error {
// 	return s.state.Save(ctx, sessionID, state)
// }

// func (s *ExplorerService) GetSchemaOverview(ctx context.Context) ([]domain.Schema, error) {
// 	rows, err := s.q.GetFullSchemaDetails(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if len(rows) == 0 {
// 		return nil, nil
// 	}

// 	var domainSchemas []domain.Schema

// 	// Start with the first schema from our sorted list.
// 	currentSchema := domain.Schema{
// 		Name:   rows[0].TableSchema,
// 		Tables: []domain.Table{},
// 	}

// 	for _, row := range rows {
// 		// Since the list is sorted by schema, a name change
// 		// means we've finished the previous schema.
// 		if row.TableSchema != currentSchema.Name {
// 			// Add the completed schema to our final slice.
// 			domainSchemas = append(domainSchemas, currentSchema)

// 			// And start the next one.
// 			currentSchema = domain.Schema{
// 				Name:   row.TableSchema,
// 				Tables: []domain.Table{},
// 			}
// 		}

// 		// Add the current table to the schema we're building.
// 		currentSchema.Tables = append(currentSchema.Tables, domain.Table{
// 			Name: row.TableName,
// 		})
// 	}

// 	// Don't forget to add the very last schema being built.
// 	domainSchemas = append(domainSchemas, currentSchema)

// 	return domainSchemas, nil
// }

// func (s *ExplorerService) GetTableColumns(ctx context.Context, schemaName, tableName string) ([]domain.Column, error) {
// 	params := store.GetTableColumnsParams{
// 		TableSchema: pgtype.Text{String: schemaName, Valid: true},
// 		TableName:   pgtype.Text{String: tableName, Valid: true},
// 	}
// 	rows, err := s.q.GetTableColumns(ctx, params)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var columns []domain.Column
// 	for _, row := range rows {
// 		columnDefault := ""
// 		if row.ColumnDefault.Valid {
// 			columnDefault = row.ColumnDefault.String
// 		}
// 		columns = append(columns, domain.Column{
// 			Name:          row.ColumnName,
// 			DataType:      row.DataType,
// 			IsNullable:    row.IsNullable,
// 			ColumnDefault: columnDefault,
// 		})
// 	}

// 	return columns, nil
// }
