package editor

// func setuproutes(router chi.router, sessionstore sessions.store, js jetstream.jetstream, pool *pgxpool.pool, q *store.queries) error {
// 	// get editor_sessions kv bucket
// 	kv, err := js.keyvalue(context.background(), "editor_sessions")
// 	if err != nil {
// 		return fmt.errorf("failed to get editor_sessions kv bucket: %w", err)
// 	}

// 	// create statestore for editorstate
// 	// statestore := session.newstatestore[services.editorstate](kv, sessionstore)

// 	// editorservice := services.neweditorservice(q, pool, statestore)
// 	// handlers := newhandlers(sessionstore, editorservice)

// 	// router.get("/editor", handlers.editorpage)

// 	// router.route("/api", func(apirouter chi.router) {
// 	// 	apirouter.route("/editor", func(editorrouter chi.router) {
// 	// 		editorrouter.get("/", handlers.editorsse)
// 	// 	})
// 	// })

// 	return nil
// }
