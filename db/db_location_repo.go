package dbora

// CONCRETE IMPLEMENTATION for domain interface
type DBLocationRepository struct {
	Conn DBTX
}

/*
func NewDBLocationRepository(db DBTX) domain.LocationRepository {
	return &DBLocationRepository{
		Conn: db,
	}
}

func (d *DBLocationRepository) Create(ctx context.Context, newLocation domain.NewLocation, newLocationAttr []domain.NewLocationAttr) error {

	var returnedId int64

	insertLocationQuery := `
	 INSERT INTO
	 	LOCATIONS (LOCATION_CODE, SAFETY_STOCK, CAPACITY, DESCRIPTION , PARENT_LOCATION_ID, TYPE_LOCATION_ID , LOCATION_STATUS, CREATED_AT , UPDATED_AT, REG_LOCATION_ID , LATITUDE, LONGITUDE, EXTERNAL_CODE, CREATED_BY, UPDATED_BY)
		VALUES   (:LOCATION_CODE, NULL        , NULL    , :DESCRIPTION, NULL              , :TYPE_LOCATION_ID, 'available'	  , CURRENT_TIMESTAMP, NULL      , :REG_LOCATION_ID, NULL    , NULL     , NULL         , NULL      , NULL)
		RETURNING ID INTO :RETURN_ID
	`

	_, err := d.Conn.ExecContext(ctx, insertLocationQuery,
		sql.Named("LOCATION_CODE", newLocation.LocationCode),
		sql.Named("DESCRIPTION", newLocation.Description),
		sql.Named("TYPE_LOCATION_ID", newLocation.TypeLocationId),
		sql.Named("REG_LOCATION_ID", newLocation.Regional),
		sql.Named("RETURN_ID", sql.Out{Dest: &returnedId}),
	)
	if err != nil {
		return err
	}

	newIdStr := strconv.FormatInt(returnedId, 10)
	log.Printf("new Id Location: %s", newIdStr)
	// resultLocationAttr, err := d.Conn.ExecContext(ctx, insertLocationAttrQuery)

	return nil
}

func (d *DBLocationRepository) FetchById(ctx context.Context, id string) (domain.LocationInfo, error) {
	var LocInfo domain.LocationInfo

	querySql := `SELECT
			L.ID,
			L.LOCATION_CODE,
			L.DESCRIPTION,
			L.TYPE_LOCATION_ID,
			L.REG_LOCATION_ID
		FROM LOCATIONS L
			WHERE L.ID = :1

	`
	err := d.Conn.QueryRowContext(ctx, querySql, id).
		Scan(
			&LocInfo.Id,
			&LocInfo.LocationCode,
			&LocInfo.Description,
			&LocInfo.TypeLocationId,
			&LocInfo.Regional,
		)
	if err == sql.ErrNoRows {
		return domain.LocationInfo{}, nil
	} else if err != nil {
		return domain.LocationInfo{}, err
	} else {
		return LocInfo, nil
	}
}

func (d *DBLocationRepository) FetchByLocationCode(ctx context.Context, locationCode string) (domain.LocationInfo, error) {
	var LocInfo domain.LocationInfo

	querySql := `SELECT
			L.ID,
			L.LOCATION_CODE,
			L.DESCRIPTION,
			L.TYPE_LOCATION_ID,
			L.REG_LOCATION_ID
		FROM LOCATIONS L
			WHERE L.LOCATION_CODE = :LOCATION_CODE

	`
	err := d.Conn.QueryRowContext(ctx, querySql, sql.Named("LOCATION_CODE", locationCode)).
		Scan(
			&LocInfo.Id,
			&LocInfo.LocationCode,
			&LocInfo.Description,
			&LocInfo.TypeLocationId,
			&LocInfo.Regional,
		)
	if err == sql.ErrNoRows {
		errMsg := fmt.Sprintf("error, lokasi dengan location_code berikut %s tidak ada di sistem", locationCode)
		return domain.LocationInfo{}, errors.New(errMsg)
	} else if err != nil {
		return domain.LocationInfo{}, err
	} else {
		return LocInfo, nil
	}
}

*/
