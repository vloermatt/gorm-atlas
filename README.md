# gorm-postgresql

A GORM managed postgresql database with Atlas integration for managing migrations 

## Schema Changes & Migration Generation

### Prerequisites:
- [Install Atlas](https://atlasgo.io/getting-started/)
- [Install Docker](https://docs.docker.com/engine/install/)


### Guide

With the prerequisites met we can now do the following

1. Update the `models.go` file to reflect desired schema changes
2. Run `atlas migrate diff --env develop` to generate a `migration.sql` file that will contain the changes made to the schema in a `sql` format
> Any potential issues with the desired change will be picked up by atlas and will result in an error 
3. Commit and push changes 

## Migration Deployment 

The `deploy-migration.yml` workflow takes care of migration application for us by using a `atlas` to apply new migrations 



