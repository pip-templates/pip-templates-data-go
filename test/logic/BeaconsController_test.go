package test_logic

// let _ = require('lodash');
// let async = require('async');
// let assert = require('chai').assert;

// import { ConfigParams } from 'pip-services3-commons-node';
// import { Descriptor } from 'pip-services3-commons-node';
// import { References } from 'pip-services3-commons-node';
// import { FilterParams } from 'pip-services3-commons-node';
// import { PagingParams } from 'pip-services3-commons-node';

// import { BeaconV1 } from '../../src/data/version1/BeaconV1';
// import { BeaconTypeV1 } from '../../src/data/version1/BeaconTypeV1';
// import { BeaconsMemoryPersistence } from '../../src/persistence/BeaconsMemoryPersistence';
// import { BeaconsController } from '../../src/logic/BeaconsController';

// const BEACON1: BeaconV1 = {
//     id: '1',
//     udi: '00001',
//     type: BeaconTypeV1.AltBeacon,
//     site_id: '1',
//     label: 'TestBeacon1',
//     center: { type: 'Point', coordinates: [ 0, 0 ] },
//     radius: 50
// };
// const BEACON2: BeaconV1 = {
//     id: '2',
//     udi: '00002',
//     type: BeaconTypeV1.iBeacon,
//     site_id: '1',
//     label: 'TestBeacon2',
//     center: { type: 'Point', coordinates: [ 2, 2 ] },
//     radius: 70
// };

// suite('BeaconsController', () => {
//     let persistence: BeaconsMemoryPersistence;
//     let controller: BeaconsController;

//     setup((done) => {
//         persistence = new BeaconsMemoryPersistence();
//         persistence.configure(new ConfigParams());

//         controller = new BeaconsController();
//         controller.configure(new ConfigParams());

//         let references = References.fromTuples(
//             new Descriptor('beacons', 'persistence', 'memory', 'default', '1.0'), persistence,
//             new Descriptor('beacons', 'controller', 'default', 'default', '1.0'), controller
//         );

//         controller.setReferences(references);

//         persistence.open(null, done);
//     });

//     teardown((done) => {
//         persistence.close(null, done);
//     });

//     test('CRUD Operations', (done) => {
//         let beacon1: BeaconV1;

//         async.series([
//             // Create the first beacon
//             (callback) => {
//                 controller.createBeacon(
//                     null,
//                     BEACON1,
//                     (err, beacon) => {
//                         assert.isNull(err);

//                         assert.isObject(beacon);
//                         assert.equal(BEACON1.udi, beacon.udi);
//                         assert.equal(BEACON1.site_id, beacon.site_id);
//                         assert.equal(BEACON1.type, beacon.type);
//                         assert.equal(BEACON1.label, beacon.label);
//                         assert.isNotNull(beacon.center);

//                         callback();
//                     }
//                 );
//             },
//             // Create the second beacon
//             (callback) => {
//                 controller.createBeacon(
//                     null,
//                     BEACON2,
//                     (err, beacon) => {
//                         assert.isNull(err);

//                         assert.isObject(beacon);
//                         assert.equal(BEACON2.udi, beacon.udi);
//                         assert.equal(BEACON2.site_id, beacon.site_id);
//                         assert.equal(BEACON2.type, beacon.type);
//                         assert.equal(BEACON2.label, beacon.label);
//                         assert.isNotNull(beacon.center);

//                         callback();
//                     }
//                 );
//             },
//             // Get all beacons
//             (callback) => {
//                 controller.getBeacons(
//                     null,
//                     new FilterParams(),
//                     new PagingParams(),
//                     (err, page) => {
//                         assert.isNull(err);

//                         assert.isObject(page);
//                         assert.lengthOf(page.data, 2);

//                         beacon1 = page.data[0];

//                         callback();
//                     }
//                 )
//             },
//             // Update the beacon
//             (callback) => {
//                 beacon1.label = 'ABC';

//                 controller.updateBeacon(
//                     null,
//                     beacon1,
//                     (err, beacon) => {
//                         assert.isNull(err);

//                         assert.isObject(beacon);
//                         assert.equal(beacon1.id, beacon.id);
//                         assert.equal('ABC', beacon.label);

//                         callback();
//                     }
//                 )
//             },
//             // Get beacon by udi
//             (callback) => {
//                 controller.getBeaconByUdi(
//                     null,
//                     beacon1.udi,
//                     (err, beacon) => {
//                         assert.isNull(err);

//                         assert.isObject(beacon);
//                         assert.equal(beacon1.id, beacon.id);

//                         callback();
//                     }
//                 )
//             },
//             // Delete the beacon
//             (callback) => {
//                 controller.deleteBeaconById(
//                     null,
//                     beacon1.id,
//                     (err, beacon) => {
//                         assert.isNull(err);

//                         assert.isObject(beacon);
//                         assert.equal(beacon1.id, beacon.id);

//                         callback();
//                     }
//                 )
//             },
//             // Try to get deleted beacon
//             (callback) => {
//                 controller.getBeaconById(
//                     null,
//                     beacon1.id,
//                     (err, beacon) => {
//                         assert.isNull(err);

//                         assert.isNull(beacon || null);

//                         callback();
//                     }
//                 )
//             }
//         ], done);
//     });

//     test('Calculate Positions', (done) => {
//         async.series([
//             // Create the first beacon
//             (callback) => {
//                 controller.createBeacon(
//                     null,
//                     BEACON1,
//                     (err, beacon) => {
//                         assert.isNull(err);

//                         assert.isObject(beacon);
//                         assert.equal(BEACON1.udi, beacon.udi);
//                         assert.equal(BEACON1.site_id, beacon.site_id);
//                         assert.equal(BEACON1.type, beacon.type);
//                         assert.equal(BEACON1.label, beacon.label);
//                         assert.isNotNull(beacon.center);

//                         callback();
//                     }
//                 );
//             },
//             // Create the second beacon
//             (callback) => {
//                 controller.createBeacon(
//                     null,
//                     BEACON2,
//                     (err, beacon) => {
//                         assert.isNull(err);

//                         assert.isObject(beacon);
//                         assert.equal(BEACON2.udi, beacon.udi);
//                         assert.equal(BEACON2.site_id, beacon.site_id);
//                         assert.equal(BEACON2.type, beacon.type);
//                         assert.equal(BEACON2.label, beacon.label);
//                         assert.isNotNull(beacon.center);

//                         callback();
//                     }
//                 );
//             },
//             // Calculate position for one beacon
//             (callback) => {
//                 controller.calculatePosition(
//                     null, '1', ['00001'],
//                     (err, position) => {
//                         assert.isNull(err);

//                         assert.isObject(position);
//                         assert.equal('Point', position.type);
//                         assert.lengthOf(position.coordinates, 2);
//                         assert.equal(0, position.coordinates[0]);
//                         assert.equal(0, position.coordinates[1]);

//                         callback();
//                     }
//                 )
//             },
//             // Calculate position for two beacons
//             (callback) => {
//                 controller.calculatePosition(
//                     null, '1', ['00001', '00002'],
//                     (err, position) => {
//                         assert.isNull(err);

//                         assert.isObject(position);
//                         assert.equal('Point', position.type);
//                         assert.lengthOf(position.coordinates, 2);
//                         assert.equal(1, position.coordinates[0]);
//                         assert.equal(1, position.coordinates[1]);

//                         callback();
//                     }
//                 )
//             }
//         ], done);
//     });
// });
