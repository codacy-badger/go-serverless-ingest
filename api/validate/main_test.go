package main_test

import (
	"testing"

	"github.com/codetaming/go-serverless-ingest/api/validate"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	headers := map[string]string{
		"describedBy": "https://schema.humancellatlas.org/type/biomaterial/5.1.0/specimen_from_organism",
	}

	tests := []struct {
		request events.APIGatewayProxyRequest
		expect  string
		err     error
	}{
		{
			// Test that the handler responds with the correct response
			// when a valid name is provided in the HTTP body

			request: events.APIGatewayProxyRequest{
				Headers: headers,
				Body: `{
    "organ": {
        "text": "brain",
        "ontology": "UBERON:0000955"
    },
    "schema_type": "biomaterial",
    "biomaterial_core": {
        "ncbi_taxon_id": [
            9606
        ],
        "biomaterial_id": "BT_S2_T",
        "has_input_biomaterial": "BT_S2",
        "biomaterial_description": "Tumor"
    },
    "organ_part": {
        "text": "temporal lobe"
    },
    "genus_species": [
        {
            "text": "Homo sapiens",
            "ontology": "NCBITaxon:9606"
        }
    ],
    "describedBy": "https://schema.humancellatlas.org/type/biomaterial/5.1.0/specimen_from_organism"
}`},
			expect: "The document is valid\n",
			err:    nil,
		},
		{
			request: events.APIGatewayProxyRequest{
				Headers: headers,
				Body:    "{}"},
			expect: "The document is not valid. see errors :\n- %s\ndescribedBy is required- %s\nschema_type is required- %s\nbiomaterial_core is required- %s\norgan is required",
			err:    nil,
		},
	}

	for _, test := range tests {
		response, err := main.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, response.Body)
	}

}
