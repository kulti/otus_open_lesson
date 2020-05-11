package router_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kulti/otus_open_lesson/internal/models"
	"github.com/kulti/otus_open_lesson/internal/router"
	"github.com/kulti/otus_open_lesson/internal/storages/memstore"
	"github.com/stretchr/testify/require"
)

func TestNotFound(t *testing.T) {
	r := router.New(nil)
	srv := httptest.NewServer(r.RootHandler())
	defer srv.Close()

	resp, err := http.Get(srv.URL + "/unknown")
	require.NoError(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestEmptyList(t *testing.T) {
	r := router.New(memstore.New())
	srv := httptest.NewServer(r.RootHandler())
	defer srv.Close()

	checkTaskList(t, srv.URL)
}

func TestCreateTask(t *testing.T) {
	r := router.New(memstore.New())
	srv := httptest.NewServer(r.RootHandler())
	defer srv.Close()

	respTask := createTask(t, srv.URL, "test task")
	checkTaskList(t, srv.URL, respTask)
}

func TestDeleteTask(t *testing.T) {
	r := router.New(memstore.New())
	srv := httptest.NewServer(r.RootHandler())
	defer srv.Close()

	respTask := createTask(t, srv.URL, "test task")
	deleteTask(t, srv.URL, respTask.ID)
	checkTaskList(t, srv.URL)
}

func checkTaskList(t *testing.T, url string, tasks ...models.Task) {
	resp, err := http.Get(url + "/tasks")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	var taskList models.TaskList
	jsDecoder := json.NewDecoder(resp.Body)
	err = jsDecoder.Decode(&taskList)
	require.NoError(t, err)

	require.Equal(t, tasks, taskList.Tasks)
}

func createTask(t *testing.T, url string, text string) models.Task {
	task := models.Task{
		Text: text,
	}
	data, err := json.Marshal(&task)
	require.NoError(t, err)

	resp, err := http.Post(url+"/task", "application/json", bytes.NewReader(data))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	var respTask models.Task
	jsDecoder := json.NewDecoder(resp.Body)
	err = jsDecoder.Decode(&respTask)
	require.NoError(t, err)

	return respTask
}

func deleteTask(t *testing.T, url string, id string) {
	req, err := http.NewRequest(http.MethodDelete, url+"/task/"+id, nil)
	require.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
}
