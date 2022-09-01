package api_key

type CreateApiKeyCommand struct {
	Name string `json:"name"`
}

func (cmd *CreateApiKeyCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(Service).Create(cmd)
}

type GetApiKeyCommand struct {
	Id string `json:"id"`
}

func (cmd *GetApiKeyCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(Service).Get(cmd)
}

type GetByApiKeyCommand struct {
	Key string `json:"key"`
}

func (cmd *GetByApiKeyCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(Service).GetByKey(cmd)
}

type ListApiKeyCommand struct {
}

func (cmd *ListApiKeyCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(Service).List(cmd)
}