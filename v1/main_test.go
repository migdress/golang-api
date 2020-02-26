package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"post-person/v1/model"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"
)

type peopleRepositoryMock struct {
	mock.Mock
}

func (m *peopleRepositoryMock) Save(p model.Person) error {
	return m.Called(p).Error(0)
}

func TestMain_adapter(t *testing.T) {

	type mocks struct {
		repo *peopleRepositoryMock
	}

	type args struct {
		r *http.Request
	}

	tests := []struct {
		name         string
		mocks        mocks
		args         args
		expectedCode int
		expectedBody string
		mocker       func(m mocks, a args)
	}{
		{
			name: "succesfully save a person",
			mocks: mocks{
				repo: &peopleRepositoryMock{},
			},
			args: args{
				r: httptest.NewRequest(
					"POST",
					"/",
					strings.NewReader(`{
						"full_name":"Ana María",
						"dni":"123",
						"birthdate":"1989-02-21"
					}`),
				),
			},
			expectedCode: http.StatusOK,
			expectedBody: ``,
			mocker: func(m mocks, a args) {
				m.repo.On(
					"Save",
					model.Person{
						FullName:  "Ana María",
						DNI:       "123",
						Birthdate: "1989-02-21",
					},
				).Return(nil).Once()
			},
		},
		{
			name: "get a 400 status code because the json is malformed",
			mocks: mocks{
				repo: &peopleRepositoryMock{},
			},
			args: args{
				r: httptest.NewRequest(
					"POST",
					"/",
					strings.NewReader(`{
						"full_name":"Ana María"
						"dni":"123",
						"birthdate":"1989-02-21"
					}`),
				),
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"error":"invalid character '"' after object key:value pair"}`,
			mocker: func(m mocks, a args) {
			},
		},
		{
			name: "get a 400 status code because the full_name field is missing",
			mocks: mocks{
				repo: &peopleRepositoryMock{},
			},
			args: args{
				r: httptest.NewRequest(
					"POST",
					"/",
					strings.NewReader(`{
						"full_name":"",
						"dni":"123",
						"birthdate":"1989-02-21"
					}`),
				),
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"errors":["full_name: String length must be greater than or equal to 1"]}`,
			mocker: func(m mocks, a args) {
			},
		},
		{
			name: "get a 500 status code because the repository returned error",
			mocks: mocks{
				repo: &peopleRepositoryMock{},
			},
			args: args{
				r: httptest.NewRequest(
					"POST",
					"/",
					strings.NewReader(`{
						"full_name":"Ana María",
						"dni":"123",
						"birthdate":"1989-02-21"
					}`),
				),
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: `{"error":"unexpected error"}`,
			mocker: func(m mocks, a args) {
				m.repo.On(
					"Save",
					model.Person{
						FullName:  "Ana María",
						DNI:       "123",
						Birthdate: "1989-02-21",
					},
				).Return(
					errors.New("unexpected error"),
				).Once()
			},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				//main()
				tt.mocker(tt.mocks, tt.args)
				handler := adapter(tt.mocks.repo)
				recorder := httptest.NewRecorder()
				handler(recorder, tt.args.r)
				if tt.expectedCode != recorder.Code {
					t.Errorf("error: expected code %v, got %v\n", tt.expectedCode, recorder.Code)
				}
				if tt.expectedBody != recorder.Body.String() {
					t.Errorf("error: expected body %v, got %v\n", tt.expectedBody, recorder.Body.String())
				}
				tt.mocks.repo.AssertExpectations(t)
			},
		)
	}
}
