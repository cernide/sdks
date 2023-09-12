/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0-rc41
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';

/**
 * The V1FailureEarlyStopping model module.
 * @module model/V1FailureEarlyStopping
 * @version 2.0.0-rc41
 */
class V1FailureEarlyStopping {
    /**
     * Constructs a new <code>V1FailureEarlyStopping</code>.
     * FailureEarlyStoppingSchema specification Early stopping based on failure config. this policy stops based on a percentage of failed runs at every evaluation.
     * @alias module:model/V1FailureEarlyStopping
     */
    constructor() {

        V1FailureEarlyStopping.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) {
    }

    /**
     * Constructs a <code>V1FailureEarlyStopping</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1FailureEarlyStopping} obj Optional instance to populate.
     * @return {module:model/V1FailureEarlyStopping} The populated <code>V1FailureEarlyStopping</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1FailureEarlyStopping();

            if (data.hasOwnProperty('kind')) {
                obj['kind'] = ApiClient.convertToType(data['kind'], 'String');
            }
            if (data.hasOwnProperty('percent')) {
                obj['percent'] = ApiClient.convertToType(data['percent'], 'Number');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>V1FailureEarlyStopping</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>V1FailureEarlyStopping</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['kind'] && !(typeof data['kind'] === 'string' || data['kind'] instanceof String)) {
            throw new Error("Expected the field `kind` to be a primitive type in the JSON string but got " + data['kind']);
        }

        return true;
    }


}



/**
 * @member {String} kind
 * @default 'failure_early_stopping'
 */
V1FailureEarlyStopping.prototype['kind'] = 'failure_early_stopping';

/**
 * The percentage of failed runs, at each evaluation interval. e.g. 1 - 99.
 * @member {Number} percent
 */
V1FailureEarlyStopping.prototype['percent'] = undefined;






export default V1FailureEarlyStopping;

