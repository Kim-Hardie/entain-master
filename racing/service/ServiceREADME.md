# Racing Service
## Overview
This service provides API methods to query racing data. Key functionalities include listing races with a variety of filters, such as by visible races or by specific meeting IDs.

### Key Adjustments
Initially, the service encountered an issue where it was unable to implement the RacingServer interface. This error originated from the service lacking certain methods mandated by the interface.

To resolve this, I implemented some necessary methods required. Specifically, the inclusion of the mustEmbedUnimplementedRacingServer() method. This method is required by newer versions of gRPC to facilitate a safe evolution of the API surface.

## Testing
### Test Framework
### Testify

- The testing framework used for this service is Testify. Testify is a flexible, easy-to-use and feature-rich testing framework for Go. We have chosen Testify because of its simplicity and resemblance to the Mockito framework in Java, which is a widely adopted testing tool.

- With Testify, it's easy to create mock objects and to assert that certain conditions hold, making it a good fit for testing our service. It simplifies the process of setting up mocks, as well as the act of verifying if the calls were made as expected.

### Test Rundown
Testing primarily covers the service's ability to list races under various filtering conditions.

#### Visible
- ShowOnlyVisible is true: This test scrutinizes the scenario where the ShowOnlyVisible filter is set to true. The expected outcome is the service's return of only the visible races.

- ShowOnlyVisible is false: This test scrutinizes the scenario where the ShowOnlyVisible filter is set to false. The expected outcome is the service's return of all races, irrespective of their visibility status.

- ShowOnlyVisible is nil: This test scrutinizes the scenario where no value is passed in for the ShowOnlyVisible filter. The expected outcome is that the service defaults to returning only the visible races.

#### Order

- OrderAscending is true: verifies that the service returns races in ascending order based on the advertised start time.

- OrderAscending is false: verifies that the service returns races in descending order based on the advertised start time.