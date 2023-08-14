/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0-rc36
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';
/**
* Enum class V1CleanPodPolicy.
* @enum {}
* @readonly
*/
export default class V1CleanPodPolicy {

        /**
         * value: "All"
         * @const
         */
        "All" = "All";


        /**
         * value: "Running"
         * @const
         */
        "Running" = "Running";


        /**
         * value: "None"
         * @const
         */
        "None" = "None";



    /**
    * Returns a <code>V1CleanPodPolicy</code> enum value from a Javascript object name.
    * @param {Object} data The plain JavaScript object containing the name of the enum value.
    * @return {module:model/V1CleanPodPolicy} The enum <code>V1CleanPodPolicy</code> value.
    */
    static constructFromObject(object) {
        return object;
    }
}

