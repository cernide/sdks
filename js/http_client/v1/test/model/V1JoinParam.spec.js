/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0-rc22
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
    instance = new PolyaxonSdk.V1JoinParam();
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

  describe('V1JoinParam', function() {
    it('should create an instance of V1JoinParam', function() {
      // uncomment below and update the code to test V1JoinParam
      //var instance = new PolyaxonSdk.V1JoinParam();
      //expect(instance).to.be.a(PolyaxonSdk.V1JoinParam);
    });

    it('should have the property value (base name: "value")', function() {
      // uncomment below and update the code to test the property value
      //var instance = new PolyaxonSdk.V1JoinParam();
      //expect(instance).to.be();
    });

    it('should have the property connection (base name: "connection")', function() {
      // uncomment below and update the code to test the property connection
      //var instance = new PolyaxonSdk.V1JoinParam();
      //expect(instance).to.be();
    });

    it('should have the property toInit (base name: "toInit")', function() {
      // uncomment below and update the code to test the property toInit
      //var instance = new PolyaxonSdk.V1JoinParam();
      //expect(instance).to.be();
    });

    it('should have the property toEnv (base name: "toEnv")', function() {
      // uncomment below and update the code to test the property toEnv
      //var instance = new PolyaxonSdk.V1JoinParam();
      //expect(instance).to.be();
    });

    it('should have the property contextOnly (base name: "contextOnly")', function() {
      // uncomment below and update the code to test the property contextOnly
      //var instance = new PolyaxonSdk.V1JoinParam();
      //expect(instance).to.be();
    });

  });

}));
