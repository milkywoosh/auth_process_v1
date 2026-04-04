package dbora

/*
func TestLocationFetchById(t *testing.T) {

	LocRepo := NewDBLocationRepository(SetupTestDBOra)

	ctxBg := context.Background()

	data, err := LocRepo.FetchById(ctxBg, "5144")
	if err != nil {
		log.Printf("err check location Id %s", err.Error())
		t.Errorf("err check location Id %s", err.Error())
	}

	log.Printf("location info %v", data)

}

func TestLocationFetchByLocationCode(t *testing.T) {

	LocRepo := NewDBLocationRepository(SetupTestDBOra)

	ctxBg := context.Background()

	expectedLocationCode := "A3713"

	data, err := LocRepo.FetchByLocationCode(ctxBg, expectedLocationCode)
	if err != nil {
		log.Printf("err check location code %s", err.Error())
		t.Errorf("err check location code %s", err.Error())
	}

	if data.LocationCode != expectedLocationCode {
		log.Printf("err value info, expect %s, got %s", expectedLocationCode, data.LocationCode)
		t.Errorf("err value info, expect %s, got %s", expectedLocationCode, data.LocationCode)
	}

	log.Printf("location info %v", data)

}


func TestCreateNewLocation(t *testing.T) {

	LocRepo := NewDBLocationRepository(SetupTestDBOra)

	ctxBg := context.Background()

	newLocationParams := domain.NewLocation{
		LocationCode:   "TEST_LOC_CODE02",
		Description:    "LOC CODE GODROR 02",
		TypeLocationId: "1",
		Regional:       "3",
	}

	err := LocRepo.Create(ctxBg, newLocationParams, []domain.NewLocationAttr{})
	if err != nil {
		log.Printf("error info insert new Location: %s", err.Error())
		t.Errorf("[FAILS] error info insert new Location: %s", err.Error())
	}

}
*/
