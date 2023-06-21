/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0-rc21
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';
import V1Optimization from './V1Optimization';

/**
 * The V1MetricEarlyStopping model module.
 * @module model/V1MetricEarlyStopping
 * @version 2.0.0-rc21
 */
class V1MetricEarlyStopping {
    /**
     * Constructs a new <code>V1MetricEarlyStopping</code>.
     * MetricEarlyStoppingSchema specification Early stopping based on metric config.
     * @alias module:model/V1MetricEarlyStopping
     */
    constructor() {

        V1MetricEarlyStopping.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) {
    }

    /**
     * Constructs a <code>V1MetricEarlyStopping</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1MetricEarlyStopping} obj Optional instance to populate.
     * @return {module:model/V1MetricEarlyStopping} The populated <code>V1MetricEarlyStopping</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1MetricEarlyStopping();

            if (data.hasOwnProperty('kind')) {
                obj['kind'] = ApiClient.convertToType(data['kind'], 'String');
            }
            if (data.hasOwnProperty('metric')) {
                obj['metric'] = ApiClient.convertToType(data['metric'], 'String');
            }
            if (data.hasOwnProperty('value')) {
                obj['value'] = ApiClient.convertToType(data['value'], 'String');
            }
            if (data.hasOwnProperty('optimization')) {
                obj['optimization'] = V1Optimization.constructFromObject(data['optimization']);
            }
            if (data.hasOwnProperty('policy')) {
                obj['policy'] = ApiClient.convertToType(data['policy'], Object);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>V1MetricEarlyStopping</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>V1MetricEarlyStopping</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['kind'] && !(typeof data['kind'] === 'string' || data['kind'] instanceof String)) {
            throw new Error("Expected the field `kind` to be a primitive type in the JSON string but got " + data['kind']);
        }
        // ensure the json data is a string
        if (data['metric'] && !(typeof data['metric'] === 'string' || data['metric'] instanceof String)) {
            throw new Error("Expected the field `metric` to be a primitive type in the JSON string but got " + data['metric']);
        }
        // ensure the json data is a string
        if (data['value'] && !(typeof data['value'] === 'string' || data['value'] instanceof String)) {
            throw new Error("Expected the field `value` to be a primitive type in the JSON string but got " + data['value']);
        }

        return true;
    }


}



/**
 * @member {String} kind
 * @default 'metric_early_stopping'
 */
V1MetricEarlyStopping.prototype['kind'] = 'metric_early_stopping';

/**
 * Metric name to use for early stopping.
 * @member {String} metric
 */
V1MetricEarlyStopping.prototype['metric'] = undefined;

/**
 * Metric value to use for the condition.
 * @member {String} value
 */
V1MetricEarlyStopping.prototype['value'] = undefined;

/**
 * @member {module:model/V1Optimization} optimization
 */
V1MetricEarlyStopping.prototype['optimization'] = undefined;

/**
 * @member {Object} policy
 */
V1MetricEarlyStopping.prototype['policy'] = undefined;






export default V1MetricEarlyStopping;

