package types

import (
	"encoding/json"
	"io"

	"github.com/senslabs/alpha/sens/errors"
	"github.com/senslabs/alpha/sens/logger"
)

func JsonMarshal(input interface{}) ([]byte, error) {
	if output, err := json.Marshal(input); err != nil {
		logger.Error(err)
		return nil, errors.New(errors.GO_ERROR, err.Error())
	} else {
		return output, nil
	}
}

func JsonMarshalToWriter(w io.Writer, input interface{}) error {
	if err := json.NewEncoder(w).Encode(input); err != nil {
		return errors.New(errors.GO_ERROR, err.Error())
	}
	return nil
}

func JsonUnmarshal(input []byte, output interface{}) error {
	if err := json.Unmarshal(input, output); err != nil {
		logger.Error(err)
		return errors.New(errors.GO_ERROR, err.Error())
	} else {
		return nil
	}
}

func JsonUnmarshalFromReader(r io.Reader, output interface{}) error {
	if err := json.NewDecoder(r).Decode(output); err != nil {
		logger.Error(err)
		return errors.New(errors.GO_ERROR, err.Error())
	}
	return nil
}

func ConvertStruct(input interface{}, output interface{}) error {
	if b, err := JsonMarshal(input); err != nil {
		return errors.New(errors.GO_ERROR, err.Error())
	} else if err := JsonUnmarshal(b, output); err != nil {
		return errors.New(errors.GO_ERROR, err.Error())
	} else {
		return nil
	}
}
