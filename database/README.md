# Database

## Entity relationship diagram 

Before creating any real database, the domain analysis was detailed in an entity relationship diagram. It can be useful to create a database using any desired relational database engine

![alt-text](/database/models/entity-relationship.png?raw=true)

# Database engine

The selected database engine was PostgreSQL, it is open source and has many useful features that most non-free engines have and can help this project a lot

# Constraints

Some stuff was taken in mind to improve the software prototype

- Every table should have a 'created-at' and 'updated-at' column
- Logical deletion should be implemented on every table, this can be done using a 'status' column
- There should be localization at the database level
- Group tables into tablespaces
- Create indexes for the most queried data

# Relational Model

Built using DBeaver

![alt-text](/database/models/relational.png?raw=true)
