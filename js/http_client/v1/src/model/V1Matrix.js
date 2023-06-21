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
import V1Bayes from './V1Bayes';
import V1GridSearch from './V1GridSearch';
import V1Hyperband from './V1Hyperband';
import V1Hyperopt from './V1Hyperopt';
import V1Iterative from './V1Iterative';
import V1Mapping from './V1Mapping';
import V1RandomSearch from './V1RandomSearch';

/**
 * The V1Matrix model module.
 * @module model/V1Matrix
 * @version 2.0.0-rc21
 */
class V1Matrix {
    /**
     * Constructs a new <code>V1Matrix</code>.
     * @alias module:model/V1Matrix
     */
    constructor() {

        V1Matrix.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) {
    }

    /**
     * Constructs a <code>V1Matrix</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1Matrix} obj Optional instance to populate.
     * @return {module:model/V1Matrix} The populated <code>V1Matrix</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1Matrix();

            if (data.hasOwnProperty('random')) {
                obj['random'] = V1RandomSearch.constructFromObject(data['random']);
            }
            if (data.hasOwnProperty('grid')) {
                obj['grid'] = V1GridSearch.constructFromObject(data['grid']);
            }
            if (data.hasOwnProperty('hyperband')) {
                obj['hyperband'] = V1Hyperband.constructFromObject(data['hyperband']);
            }
            if (data.hasOwnProperty('bayes')) {
                obj['bayes'] = V1Bayes.constructFromObject(data['bayes']);
            }
            if (data.hasOwnProperty('hyperopt')) {
                obj['hyperopt'] = V1Hyperopt.constructFromObject(data['hyperopt']);
            }
            if (data.hasOwnProperty('iterative')) {
                obj['iterative'] = V1Iterative.constructFromObject(data['iterative']);
            }
            if (data.hasOwnProperty('mapping')) {
                obj['mapping'] = V1Mapping.constructFromObject(data['mapping']);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>V1Matrix</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>V1Matrix</code>.
     */
    static validateJSON(data) {
        // validate the optional field `random`
        if (data['random']) { // data not null
          V1RandomSearch.validateJSON(data['random']);
        }
        // validate the optional field `grid`
        if (data['grid']) { // data not null
          V1GridSearch.validateJSON(data['grid']);
        }
        // validate the optional field `hyperband`
        if (data['hyperband']) { // data not null
          V1Hyperband.validateJSON(data['hyperband']);
        }
        // validate the optional field `bayes`
        if (data['bayes']) { // data not null
          V1Bayes.validateJSON(data['bayes']);
        }
        // validate the optional field `hyperopt`
        if (data['hyperopt']) { // data not null
          V1Hyperopt.validateJSON(data['hyperopt']);
        }
        // validate the optional field `iterative`
        if (data['iterative']) { // data not null
          V1Iterative.validateJSON(data['iterative']);
        }
        // validate the optional field `mapping`
        if (data['mapping']) { // data not null
          V1Mapping.validateJSON(data['mapping']);
        }

        return true;
    }


}



/**
 * @member {module:model/V1RandomSearch} random
 */
V1Matrix.prototype['random'] = undefined;

/**
 * @member {module:model/V1GridSearch} grid
 */
V1Matrix.prototype['grid'] = undefined;

/**
 * @member {module:model/V1Hyperband} hyperband
 */
V1Matrix.prototype['hyperband'] = undefined;

/**
 * @member {module:model/V1Bayes} bayes
 */
V1Matrix.prototype['bayes'] = undefined;

/**
 * @member {module:model/V1Hyperopt} hyperopt
 */
V1Matrix.prototype['hyperopt'] = undefined;

/**
 * @member {module:model/V1Iterative} iterative
 */
V1Matrix.prototype['iterative'] = undefined;

/**
 * @member {module:model/V1Mapping} mapping
 */
V1Matrix.prototype['mapping'] = undefined;






export default V1Matrix;

