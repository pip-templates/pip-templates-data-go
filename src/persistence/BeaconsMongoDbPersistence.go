let _ = require('lodash');

import { FilterParams } from 'pip-services3-commons-node';
import { PagingParams } from 'pip-services3-commons-node';
import { DataPage } from 'pip-services3-commons-node';

import { IdentifiableMongoDbPersistence } from 'pip-services3-mongodb-node';

import { BeaconV1 } from '../data/version1/BeaconV1';
import { IBeaconsPersistence } from './IBeaconsPersistence';
import { BeaconsMongoDbSchema } from './BeaconsMongoDbSchema';

export class BeaconsMongoDbPersistence
    extends IdentifiableMongoDbPersistence<BeaconV1, string>
    implements IBeaconsPersistence {
    constructor() {
        super('beacons', BeaconsMongoDbSchema());
        this._maxPageSize = 1000;
    }

    private composeFilter(filter: FilterParams): any {
        filter = filter || new FilterParams();

        let criteria = [];

        let id = filter.getAsNullableString('id');
        if (id != null) 
            criteria.push({ _id: id });

        let siteId = filter.getAsNullableString('site_id');
        if (siteId != null)
            criteria.push({ site_id: siteId });

        let label = filter.getAsNullableString('label');
        if (label != null)
            criteria.push({ label: label });

        let udi = filter.getAsNullableString('udi');
        if (udi != null) {
            criteria.push({ udi: udi });
        }

        let udis = filter.getAsObject('udis');
        if (_.isString(udis))
            udis = udis.split(',');
        if (_.isArray(udis))
            criteria.push({ udi: { $in: udis } });

        return criteria.length > 0 ? { $and: criteria } : null;
    }

    public getPageByFilter(correlationId: string, filter: FilterParams, paging: PagingParams,
        callback: (err: any, page: DataPage<BeaconV1>) => void): void {
        super.getPageByFilter(correlationId, this.composeFilter(filter), paging, null, null, callback);
    }

    public getOneByUdi(correlationId: string, udi: string,
        callback: (err: any, item: BeaconV1) => void): void {

        let criteria = {
            udi: udi
        };

        this._model.findOne(criteria, (err, item) => {
            item = this.convertFromPublic(item);

            if (item != null) this._logger.trace(correlationId, "Found beacon by %s", udi);
            else this._logger.trace(correlationId, "Cannot find beacon by %s", udi);

            callback(err, item);
        });
    }
}