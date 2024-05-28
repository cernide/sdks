/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.2.0
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';
/**
* Enum class V1RunKind.
* @enum {}
* @readonly
*/
export default class V1RunKind {

        /**
         * value: "job"
         * @const
         */
        "job" = "job";


        /**
         * value: "service"
         * @const
         */
        "service" = "service";


        /**
         * value: "dag"
         * @const
         */
        "dag" = "dag";


        /**
         * value: "daskjob"
         * @const
         */
        "daskjob" = "daskjob";


        /**
         * value: "rayjob"
         * @const
         */
        "rayjob" = "rayjob";


        /**
         * value: "mpijob"
         * @const
         */
        "mpijob" = "mpijob";


        /**
         * value: "tfjob"
         * @const
         */
        "tfjob" = "tfjob";


        /**
         * value: "pytorchjob"
         * @const
         */
        "pytorchjob" = "pytorchjob";


        /**
         * value: "mxjob"
         * @const
         */
        "mxjob" = "mxjob";


        /**
         * value: "xgbjob"
         * @const
         */
        "xgbjob" = "xgbjob";


        /**
         * value: "paddlejob"
         * @const
         */
        "paddlejob" = "paddlejob";


        /**
         * value: "matrix"
         * @const
         */
        "matrix" = "matrix";


        /**
         * value: "schedule"
         * @const
         */
        "schedule" = "schedule";


        /**
         * value: "tuner"
         * @const
         */
        "tuner" = "tuner";


        /**
         * value: "watchdog"
         * @const
         */
        "watchdog" = "watchdog";


        /**
         * value: "notifier"
         * @const
         */
        "notifier" = "notifier";


        /**
         * value: "builder"
         * @const
         */
        "builder" = "builder";


        /**
         * value: "cleaner"
         * @const
         */
        "cleaner" = "cleaner";



    /**
    * Returns a <code>V1RunKind</code> enum value from a Javascript object name.
    * @param {Object} data The plain JavaScript object containing the name of the enum value.
    * @return {module:model/V1RunKind} The enum <code>V1RunKind</code> value.
    */
    static constructFromObject(object) {
        return object;
    }
}

