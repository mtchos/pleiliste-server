// Code generated by MockGen. DO NOT EDIT.
// Source: internal/adapter/integration/spotify/spotify_integration.go
//
// Generated by this command:
//
//	mockgen --source=internal/adapter/integration/spotify/spotify_integration.go --destination=internal/adapter/integration/spotify/mock/mock_spotify_client.go
//

// Package mock_spotify is a generated GoMock package.
package mock_spotify

import (
	reflect "reflect"

	spotify "github.com/mtchos/pleiback/internal/adapter/integration/spotify"
	entity "github.com/mtchos/pleiback/internal/domain/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockIntegration is a mock of Integration interface.
type MockIntegration struct {
	ctrl     *gomock.Controller
	recorder *MockIntegrationMockRecorder
	isgomock struct{}
}

// MockIntegrationMockRecorder is the mock recorder for MockIntegration.
type MockIntegrationMockRecorder struct {
	mock *MockIntegration
}

// NewMockIntegration creates a new mock instance.
func NewMockIntegration(ctrl *gomock.Controller) *MockIntegration {
	mock := &MockIntegration{ctrl: ctrl}
	mock.recorder = &MockIntegrationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIntegration) EXPECT() *MockIntegrationMockRecorder {
	return m.recorder
}

// AddPlaylistTracks mocks base method.
func (m *MockIntegration) AddPlaylistTracks(playlistID string, tracksIDs []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPlaylistTracks", playlistID, tracksIDs)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPlaylistTracks indicates an expected call of AddPlaylistTracks.
func (mr *MockIntegrationMockRecorder) AddPlaylistTracks(playlistID, tracksIDs any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPlaylistTracks", reflect.TypeOf((*MockIntegration)(nil).AddPlaylistTracks), playlistID, tracksIDs)
}

// CreatePlaylist mocks base method.
func (m *MockIntegration) CreatePlaylist(userID string, playlist entity.Playlist) (spotify.CreatePlaylistResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePlaylist", userID, playlist)
	ret0, _ := ret[0].(spotify.CreatePlaylistResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePlaylist indicates an expected call of CreatePlaylist.
func (mr *MockIntegrationMockRecorder) CreatePlaylist(userID, playlist any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePlaylist", reflect.TypeOf((*MockIntegration)(nil).CreatePlaylist), userID, playlist)
}

// GetArtists mocks base method.
func (m *MockIntegration) GetArtists(artistsIDs []string) (spotify.GetArtistsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArtists", artistsIDs)
	ret0, _ := ret[0].(spotify.GetArtistsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArtists indicates an expected call of GetArtists.
func (mr *MockIntegrationMockRecorder) GetArtists(artistsIDs any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArtists", reflect.TypeOf((*MockIntegration)(nil).GetArtists), artistsIDs)
}

// SearchTracks mocks base method.
func (m *MockIntegration) SearchTracks(query string, limit, offset int64) (spotify.SearchTracksResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchTracks", query, limit, offset)
	ret0, _ := ret[0].(spotify.SearchTracksResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchTracks indicates an expected call of SearchTracks.
func (mr *MockIntegrationMockRecorder) SearchTracks(query, limit, offset any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchTracks", reflect.TypeOf((*MockIntegration)(nil).SearchTracks), query, limit, offset)
}
