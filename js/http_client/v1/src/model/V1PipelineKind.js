/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.1.0-rc5
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';
/**
* Enum class V1PipelineKind.
* @enum {}
* @readonly
*/
export default class V1PipelineKind {

        /**
         * value: "dag"
         * @const
         */
        "dag" = "dag";


        /**
         * value: "matrix"
         * @const
         */
        "matrix" = "matrix";



    /**
    * Returns a <code>V1PipelineKind</code> enum value from a Javascript object name.
    * @param {Object} data The plain JavaScript object containing the name of the enum value.
    * @return {module:model/V1PipelineKind} The enum <code>V1PipelineKind</code> value.
    */
    static constructFromObject(object) {
        return object;
    }
}

