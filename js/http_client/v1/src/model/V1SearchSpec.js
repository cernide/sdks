/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0-rc33
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';
import V1AnalyticsSpec from './V1AnalyticsSpec';
import V1DashboardSpec from './V1DashboardSpec';

/**
 * The V1SearchSpec model module.
 * @module model/V1SearchSpec
 * @version 2.0.0-rc33
 */
class V1SearchSpec {
    /**
     * Constructs a new <code>V1SearchSpec</code>.
     * @alias module:model/V1SearchSpec
     */
    constructor() {

        V1SearchSpec.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) {
    }

    /**
     * Constructs a <code>V1SearchSpec</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1SearchSpec} obj Optional instance to populate.
     * @return {module:model/V1SearchSpec} The populated <code>V1SearchSpec</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1SearchSpec();

            if (data.hasOwnProperty('query')) {
                obj['query'] = ApiClient.convertToType(data['query'], 'String');
            }
            if (data.hasOwnProperty('sort')) {
                obj['sort'] = ApiClient.convertToType(data['sort'], 'String');
            }
            if (data.hasOwnProperty('limit')) {
                obj['limit'] = ApiClient.convertToType(data['limit'], 'Number');
            }
            if (data.hasOwnProperty('offset')) {
                obj['offset'] = ApiClient.convertToType(data['offset'], 'Number');
            }
            if (data.hasOwnProperty('groupby')) {
                obj['groupby'] = ApiClient.convertToType(data['groupby'], 'String');
            }
            if (data.hasOwnProperty('columns')) {
                obj['columns'] = ApiClient.convertToType(data['columns'], 'String');
            }
            if (data.hasOwnProperty('layout')) {
                obj['layout'] = ApiClient.convertToType(data['layout'], 'String');
            }
            if (data.hasOwnProperty('sections')) {
                obj['sections'] = ApiClient.convertToType(data['sections'], 'String');
            }
            if (data.hasOwnProperty('compares')) {
                obj['compares'] = ApiClient.convertToType(data['compares'], 'String');
            }
            if (data.hasOwnProperty('heat')) {
                obj['heat'] = ApiClient.convertToType(data['heat'], 'String');
            }
            if (data.hasOwnProperty('events')) {
                obj['events'] = V1DashboardSpec.constructFromObject(data['events']);
            }
            if (data.hasOwnProperty('histograms')) {
                obj['histograms'] = ApiClient.convertToType(data['histograms'], Object);
            }
            if (data.hasOwnProperty('trends')) {
                obj['trends'] = ApiClient.convertToType(data['trends'], Object);
            }
            if (data.hasOwnProperty('analytics')) {
                obj['analytics'] = V1AnalyticsSpec.constructFromObject(data['analytics']);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>V1SearchSpec</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>V1SearchSpec</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['query'] && !(typeof data['query'] === 'string' || data['query'] instanceof String)) {
            throw new Error("Expected the field `query` to be a primitive type in the JSON string but got " + data['query']);
        }
        // ensure the json data is a string
        if (data['sort'] && !(typeof data['sort'] === 'string' || data['sort'] instanceof String)) {
            throw new Error("Expected the field `sort` to be a primitive type in the JSON string but got " + data['sort']);
        }
        // ensure the json data is a string
        if (data['groupby'] && !(typeof data['groupby'] === 'string' || data['groupby'] instanceof String)) {
            throw new Error("Expected the field `groupby` to be a primitive type in the JSON string but got " + data['groupby']);
        }
        // ensure the json data is a string
        if (data['columns'] && !(typeof data['columns'] === 'string' || data['columns'] instanceof String)) {
            throw new Error("Expected the field `columns` to be a primitive type in the JSON string but got " + data['columns']);
        }
        // ensure the json data is a string
        if (data['layout'] && !(typeof data['layout'] === 'string' || data['layout'] instanceof String)) {
            throw new Error("Expected the field `layout` to be a primitive type in the JSON string but got " + data['layout']);
        }
        // ensure the json data is a string
        if (data['sections'] && !(typeof data['sections'] === 'string' || data['sections'] instanceof String)) {
            throw new Error("Expected the field `sections` to be a primitive type in the JSON string but got " + data['sections']);
        }
        // ensure the json data is a string
        if (data['compares'] && !(typeof data['compares'] === 'string' || data['compares'] instanceof String)) {
            throw new Error("Expected the field `compares` to be a primitive type in the JSON string but got " + data['compares']);
        }
        // ensure the json data is a string
        if (data['heat'] && !(typeof data['heat'] === 'string' || data['heat'] instanceof String)) {
            throw new Error("Expected the field `heat` to be a primitive type in the JSON string but got " + data['heat']);
        }
        // validate the optional field `events`
        if (data['events']) { // data not null
          V1DashboardSpec.validateJSON(data['events']);
        }
        // validate the optional field `analytics`
        if (data['analytics']) { // data not null
          V1AnalyticsSpec.validateJSON(data['analytics']);
        }

        return true;
    }


}



/**
 * @member {String} query
 */
V1SearchSpec.prototype['query'] = undefined;

/**
 * @member {String} sort
 */
V1SearchSpec.prototype['sort'] = undefined;

/**
 * @member {Number} limit
 */
V1SearchSpec.prototype['limit'] = undefined;

/**
 * @member {Number} offset
 */
V1SearchSpec.prototype['offset'] = undefined;

/**
 * @member {String} groupby
 */
V1SearchSpec.prototype['groupby'] = undefined;

/**
 * @member {String} columns
 */
V1SearchSpec.prototype['columns'] = undefined;

/**
 * @member {String} layout
 */
V1SearchSpec.prototype['layout'] = undefined;

/**
 * @member {String} sections
 */
V1SearchSpec.prototype['sections'] = undefined;

/**
 * @member {String} compares
 */
V1SearchSpec.prototype['compares'] = undefined;

/**
 * @member {String} heat
 */
V1SearchSpec.prototype['heat'] = undefined;

/**
 * @member {module:model/V1DashboardSpec} events
 */
V1SearchSpec.prototype['events'] = undefined;

/**
 * @member {Object} histograms
 */
V1SearchSpec.prototype['histograms'] = undefined;

/**
 * @member {Object} trends
 */
V1SearchSpec.prototype['trends'] = undefined;

/**
 * @member {module:model/V1AnalyticsSpec} analytics
 */
V1SearchSpec.prototype['analytics'] = undefined;






export default V1SearchSpec;

