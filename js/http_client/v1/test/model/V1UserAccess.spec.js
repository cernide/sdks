/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.3.3
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
    instance = new PolyaxonSdk.V1UserAccess();
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

  describe('V1UserAccess', function() {
    it('should create an instance of V1UserAccess', function() {
      // uncomment below and update the code to test V1UserAccess
      //var instance = new PolyaxonSdk.V1UserAccess();
      //expect(instance).to.be.a(PolyaxonSdk.V1UserAccess);
    });

    it('should have the property user (base name: "user")', function() {
      // uncomment below and update the code to test the property user
      //var instance = new PolyaxonSdk.V1UserAccess();
      //expect(instance).to.be();
    });

    it('should have the property user_data (base name: "user_data")', function() {
      // uncomment below and update the code to test the property user_data
      //var instance = new PolyaxonSdk.V1UserAccess();
      //expect(instance).to.be();
    });

    it('should have the property queue (base name: "queue")', function() {
      // uncomment below and update the code to test the property queue
      //var instance = new PolyaxonSdk.V1UserAccess();
      //expect(instance).to.be();
    });

    it('should have the property preset (base name: "preset")', function() {
      // uncomment below and update the code to test the property preset
      //var instance = new PolyaxonSdk.V1UserAccess();
      //expect(instance).to.be();
    });

    it('should have the property namespace (base name: "namespace")', function() {
      // uncomment below and update the code to test the property namespace
      //var instance = new PolyaxonSdk.V1UserAccess();
      //expect(instance).to.be();
    });

  });

}));
