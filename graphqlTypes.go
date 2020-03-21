package main

import (
	"github.com/graphql-go/graphql"
	"github.com/lruggieri/graphqltest/pkg/company"
	"github.com/lruggieri/graphqltest/pkg/database"
	"github.com/lruggieri/graphqltest/pkg/job"
)

// jobType defines the Job type for GraphQL
var jobType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Job",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"position": &graphql.Field{
				Type: graphql.String,
			},
			"companyID": &graphql.Field{
				Type: graphql.Int,
			},
			"company": &graphql.Field{
				Type: companyType,
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					companyID := params.Source.(job.Job).CompanyID
					return database.GetCompanyByID(companyID), nil
				},
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"location": &graphql.Field{
				Type: graphql.String,
			},
			"employmentType": &graphql.Field{
				Type: graphql.String,
			},
			"skillsRequired": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
		},
	},
)

// companyType defines the Company type for GraphQL
var companyType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Company",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"ceoID": &graphql.Field{
				Type: graphql.Int,
			},
			"ceo": &graphql.Field{
				Type: personType,
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					ceoID := params.Source.(company.Company).CeoID
					return database.GetPersonByID(ceoID), nil
				},
			},
		},
	},
)

// personType defines the Person type for GraphQL
var personType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Person",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"age": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)
