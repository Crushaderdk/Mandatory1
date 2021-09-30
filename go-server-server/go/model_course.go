/*
 * Swagger Studentcourse
 *
 * This is a sample server Studentcourse server.  You can find out more about     Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).      For this sample, you can use the api key `special-key` to test the authorization     filters.
 *
 * API version: 1.0.0
 * Contact: apiteam@swagger.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Course struct {
	Id int64 `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Students []interface{} `json:"students,omitempty"`

	Teachers []interface{} `json:"teachers,omitempty"`

	Workload float64 `json:"workload,omitempty"`

	StudentSatisfactionRating float64 `json:"studentSatisfactionRating,omitempty"`
}