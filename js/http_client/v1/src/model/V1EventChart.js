/**
 * Polyaxon SDKs and REST API specification.
 *    
 *
 * The version of the OpenAPI document: 2.1.8
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';
import V1EventChartKind from './V1EventChartKind';

/**
 * The V1EventChart model module.
 * @module model/V1EventChart
 * @version 2.1.8
 */
class V1EventChart {
    /**
     * Constructs a new <code>V1EventChart</code>.
     * @alias module:model/V1EventChart
     */
    constructor() { 
        
        V1EventChart.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1EventChart</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1EventChart} obj Optional instance to populate.
     * @return {module:model/V1EventChart} The populated <code>V1EventChart</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1EventChart();

            if (data.hasOwnProperty('kind')) {
                obj['kind'] = V1EventChartKind.constructFromObject(data['kind']);
            }
            if (data.hasOwnProperty('figure')) {
                obj['figure'] = ApiClient.convertToType(data['figure'], Object);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>V1EventChart</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>V1EventChart</code>.
     */
    static validateJSON(data) {

        return true;
    }


}



/**
 * @member {module:model/V1EventChartKind} kind
 */
V1EventChart.prototype['kind'] = undefined;

/**
 * @member {Object} figure
 */
V1EventChart.prototype['figure'] = undefined;






export default V1EventChart;

