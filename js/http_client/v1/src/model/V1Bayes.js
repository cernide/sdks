/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0-rc28
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';
import V1OptimizationMetric from './V1OptimizationMetric';
import V1Tuner from './V1Tuner';

/**
 * The V1Bayes model module.
 * @module model/V1Bayes
 * @version 2.0.0-rc28
 */
class V1Bayes {
    /**
     * Constructs a new <code>V1Bayes</code>.
     * @alias module:model/V1Bayes
     */
    constructor() {

        V1Bayes.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) {
    }

    /**
     * Constructs a <code>V1Bayes</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1Bayes} obj Optional instance to populate.
     * @return {module:model/V1Bayes} The populated <code>V1Bayes</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1Bayes();

            if (data.hasOwnProperty('kind')) {
                obj['kind'] = ApiClient.convertToType(data['kind'], 'String');
            }
            if (data.hasOwnProperty('params')) {
                obj['params'] = ApiClient.convertToType(data['params'], {'String': Object});
            }
            if (data.hasOwnProperty('numInitialRuns')) {
                obj['numInitialRuns'] = ApiClient.convertToType(data['numInitialRuns'], 'Number');
            }
            if (data.hasOwnProperty('maxIterations')) {
                obj['maxIterations'] = ApiClient.convertToType(data['maxIterations'], 'Number');
            }
            if (data.hasOwnProperty('utilityFunction')) {
                obj['utilityFunction'] = ApiClient.convertToType(data['utilityFunction'], Object);
            }
            if (data.hasOwnProperty('metric')) {
                obj['metric'] = V1OptimizationMetric.constructFromObject(data['metric']);
            }
            if (data.hasOwnProperty('seed')) {
                obj['seed'] = ApiClient.convertToType(data['seed'], 'Number');
            }
            if (data.hasOwnProperty('concurrency')) {
                obj['concurrency'] = ApiClient.convertToType(data['concurrency'], 'Number');
            }
            if (data.hasOwnProperty('tuner')) {
                obj['tuner'] = V1Tuner.constructFromObject(data['tuner']);
            }
            if (data.hasOwnProperty('earlyStopping')) {
                obj['earlyStopping'] = ApiClient.convertToType(data['earlyStopping'], [Object]);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>V1Bayes</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>V1Bayes</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['kind'] && !(typeof data['kind'] === 'string' || data['kind'] instanceof String)) {
            throw new Error("Expected the field `kind` to be a primitive type in the JSON string but got " + data['kind']);
        }
        // validate the optional field `metric`
        if (data['metric']) { // data not null
          V1OptimizationMetric.validateJSON(data['metric']);
        }
        // validate the optional field `tuner`
        if (data['tuner']) { // data not null
          V1Tuner.validateJSON(data['tuner']);
        }
        // ensure the json data is an array
        if (!Array.isArray(data['earlyStopping'])) {
            throw new Error("Expected the field `earlyStopping` to be an array in the JSON data but got " + data['earlyStopping']);
        }

        return true;
    }


}



/**
 * @member {String} kind
 * @default 'bayes'
 */
V1Bayes.prototype['kind'] = 'bayes';

/**
 * @member {Object.<String, Object>} params
 */
V1Bayes.prototype['params'] = undefined;

/**
 * @member {Number} numInitialRuns
 */
V1Bayes.prototype['numInitialRuns'] = undefined;

/**
 * @member {Number} maxIterations
 */
V1Bayes.prototype['maxIterations'] = undefined;

/**
 * @member {Object} utilityFunction
 */
V1Bayes.prototype['utilityFunction'] = undefined;

/**
 * @member {module:model/V1OptimizationMetric} metric
 */
V1Bayes.prototype['metric'] = undefined;

/**
 * @member {Number} seed
 */
V1Bayes.prototype['seed'] = undefined;

/**
 * @member {Number} concurrency
 */
V1Bayes.prototype['concurrency'] = undefined;

/**
 * @member {module:model/V1Tuner} tuner
 */
V1Bayes.prototype['tuner'] = undefined;

/**
 * @member {Array.<Object>} earlyStopping
 */
V1Bayes.prototype['earlyStopping'] = undefined;






export default V1Bayes;

