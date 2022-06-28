package user

// func TestRepositoryGet_Ok(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	assert.NoError(t, err)
// 	defer db.Close()

// 	ctx := context.Background()
// 	repository := NewRepository(db)

// 	rows := sqlmock.NewRows([]string{
// 		"id",
// 		"ldap_user",
// 		"first_name",
// 		"last_name",
// 		"email",
// 		"meli_site",
// 	}).AddRow(
// 		"ID",
// 		"LDAP",
// 		"FirstName",
// 		"LastName",
// 		"email@mercadolibre.com",
// 		"site",
// 	)
// 	mock.ExpectQuery("SELECT").WillReturnRows(rows)
// 	get, err := repository.Get(ctx, "ID")
// 	assert.NoError(t, err)
// 	assert.Equal(t, "FirstName", get.FirstName)

// }
