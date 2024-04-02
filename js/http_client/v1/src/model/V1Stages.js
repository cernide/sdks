/**
 * Polyaxon SDKs and REST API specification.
 *    
 *
 * The version of the OpenAPI document: 2.1.5
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';
/**
* Enum class V1Stages.
* @enum {}
* @readonly
*/
export default class V1Stages {
    
        /**
         * value: "testing"
         * @const
         */
        "testing" = "testing";

    
        /**
         * value: "staging"
         * @const
         */
        "staging" = "staging";

    
        /**
         * value: "production"
         * @const
         */
        "production" = "production";

    
        /**
         * value: "disabled"
         * @const
         */
        "disabled" = "disabled";

    

    /**
    * Returns a <code>V1Stages</code> enum value from a Javascript object name.
    * @param {Object} data The plain JavaScript object containing the name of the enum value.
    * @return {module:model/V1Stages} The enum <code>V1Stages</code> value.
    */
    static constructFromObject(object) {
        return object;
    }
}

