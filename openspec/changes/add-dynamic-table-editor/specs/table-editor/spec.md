# Table Editor Specification

## ADDED Requirements

### Requirement: Table Selection
The editor SHALL provide a simple mechanism for users to select which database table to view.

#### Scenario: User selects a table using text input
- **GIVEN** the editor page is loaded
- **WHEN** the user enters a table identifier in the format "schema.table" (e.g., "public.users")
- **THEN** the system SHALL parse the schema and table name
- **AND** the system SHALL load and display the data for that table

#### Scenario: Invalid table name handling
- **GIVEN** the user enters an invalid or non-existent table name
- **WHEN** the system attempts to load the table data
- **THEN** the system SHALL display an error message indicating the table was not found
- **AND** the previous table data (if any) SHALL remain visible

### Requirement: Dynamic Table Data Display
The editor SHALL fetch and display actual table data from the selected PostgreSQL table.

#### Scenario: Display table data with correct columns
- **GIVEN** a table is selected
- **WHEN** the table data is loaded
- **THEN** the system SHALL query the table's column information from information_schema
- **AND** the system SHALL display column names as table headers
- **AND** the system SHALL display row data matching the column order

#### Scenario: Handle empty tables
- **GIVEN** a table with no rows is selected
- **WHEN** the table data is loaded
- **THEN** the system SHALL display the column headers
- **AND** the system SHALL display a message indicating the table is empty

### Requirement: Pagination Support
The editor SHALL provide pagination controls to navigate through large datasets.

#### Scenario: Display pagination controls
- **GIVEN** a table with data is loaded
- **WHEN** the table view is rendered
- **THEN** the system SHALL display pagination controls at the bottom of the table
- **AND** the controls SHALL include: previous page button, next page button, and current page indicator

#### Scenario: Navigate to next page
- **GIVEN** the user is viewing page 1 of table data
- **AND** there are more rows available
- **WHEN** the user clicks the "next page" button
- **THEN** the system SHALL load and display the next page of data
- **AND** the page indicator SHALL update to show page 2

#### Scenario: Navigate to previous page
- **GIVEN** the user is viewing page 2 or higher
- **WHEN** the user clicks the "previous page" button
- **THEN** the system SHALL load and display the previous page of data
- **AND** the page indicator SHALL decrement

#### Scenario: Disable navigation at boundaries
- **GIVEN** the user is on the first page
- **THEN** the previous page button SHALL be disabled
- **GIVEN** the user is on the last page
- **THEN** the next page button SHALL be disabled

### Requirement: Page Size Configuration
The editor SHALL use a fixed page size for pagination.

#### Scenario: Default page size
- **GIVEN** any table is loaded with pagination
- **WHEN** data is fetched
- **THEN** the system SHALL retrieve and display 50 rows per page by default

### Requirement: Data Query Performance
The editor SHALL use efficient SQL queries with LIMIT and OFFSET for pagination.

#### Scenario: Fetch paginated data
- **GIVEN** a table and page number are specified
- **WHEN** the system queries the database
- **THEN** the query SHALL use LIMIT to restrict the number of rows
- **AND** the query SHALL use OFFSET to skip rows from previous pages
- **AND** the query SHALL preserve the table's natural order (or order by primary key if available)

#### Scenario: Count total rows for pagination
- **GIVEN** a table is selected
- **WHEN** pagination is initialized
- **THEN** the system SHALL execute a COUNT(*) query to determine total rows
- **AND** the system SHALL calculate the total number of pages based on page size
