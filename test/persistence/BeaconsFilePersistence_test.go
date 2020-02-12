import { ConfigParams } from 'pip-services3-commons-node';

import { BeaconsFilePersistence } from '../../src/persistence/BeaconsFilePersistence';
import { BeaconsPersistenceFixture } from './BeaconsPersistenceFixture';

suite('BeaconsFilePersistence', () => {
    let persistence: BeaconsFilePersistence;
    let fixture: BeaconsPersistenceFixture;

    setup((done) => {
        persistence = new BeaconsFilePersistence('data/beacons.test.json');
        persistence.configure(new ConfigParams());

        fixture = new BeaconsPersistenceFixture(persistence);

        persistence.open(null, (err) => {
            persistence.clear(null, done);
        });
    });

    teardown((done) => {
        persistence.close(null, done);
    });

    test('CRUD Operations', (done) => {
        fixture.testCrudOperations(done);
    });

    test('Get with Filters', (done) => {
        fixture.testGetWithFilters(done);
    });

});