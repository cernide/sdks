/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.1.2-rc1
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', process.cwd()+'/src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require(process.cwd()+'/src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.PolyaxonSdk);
  }
}(this, function(expect, PolyaxonSdk) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new PolyaxonSdk.V1StageCondition();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('V1StageCondition', function() {
    it('should create an instance of V1StageCondition', function() {
      // uncomment below and update the code to test V1StageCondition
      //var instance = new PolyaxonSdk.V1StageCondition();
      //expect(instance).to.be.a(PolyaxonSdk.V1StageCondition);
    });

    it('should have the property type (base name: "type")', function() {
      // uncomment below and update the code to test the property type
      //var instance = new PolyaxonSdk.V1StageCondition();
      //expect(instance).to.be();
    });

    it('should have the property status (base name: "status")', function() {
      // uncomment below and update the code to test the property status
      //var instance = new PolyaxonSdk.V1StageCondition();
      //expect(instance).to.be();
    });

    it('should have the property reason (base name: "reason")', function() {
      // uncomment below and update the code to test the property reason
      //var instance = new PolyaxonSdk.V1StageCondition();
      //expect(instance).to.be();
    });

    it('should have the property message (base name: "message")', function() {
      // uncomment below and update the code to test the property message
      //var instance = new PolyaxonSdk.V1StageCondition();
      //expect(instance).to.be();
    });

    it('should have the property last_update_time (base name: "last_update_time")', function() {
      // uncomment below and update the code to test the property last_update_time
      //var instance = new PolyaxonSdk.V1StageCondition();
      //expect(instance).to.be();
    });

    it('should have the property last_transition_time (base name: "last_transition_time")', function() {
      // uncomment below and update the code to test the property last_transition_time
      //var instance = new PolyaxonSdk.V1StageCondition();
      //expect(instance).to.be();
    });

    it('should have the property meta_info (base name: "meta_info")', function() {
      // uncomment below and update the code to test the property meta_info
      //var instance = new PolyaxonSdk.V1StageCondition();
      //expect(instance).to.be();
    });

  });

}));
