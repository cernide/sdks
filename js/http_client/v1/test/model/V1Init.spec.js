/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0
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
    instance = new PolyaxonSdk.V1Init();
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

  describe('V1Init', function() {
    it('should create an instance of V1Init', function() {
      // uncomment below and update the code to test V1Init
      //var instance = new PolyaxonSdk.V1Init();
      //expect(instance).to.be.a(PolyaxonSdk.V1Init);
    });

    it('should have the property artifacts (base name: "artifacts")', function() {
      // uncomment below and update the code to test the property artifacts
      //var instance = new PolyaxonSdk.V1Init();
      //expect(instance).to.be();
    });

    it('should have the property paths (base name: "paths")', function() {
      // uncomment below and update the code to test the property paths
      //var instance = new PolyaxonSdk.V1Init();
      //expect(instance).to.be();
    });

    it('should have the property git (base name: "git")', function() {
      // uncomment below and update the code to test the property git
      //var instance = new PolyaxonSdk.V1Init();
      //expect(instance).to.be();
    });

    it('should have the property dockerfile (base name: "dockerfile")', function() {
      // uncomment below and update the code to test the property dockerfile
      //var instance = new PolyaxonSdk.V1Init();
      //expect(instance).to.be();
    });

    it('should have the property file (base name: "file")', function() {
      // uncomment below and update the code to test the property file
      //var instance = new PolyaxonSdk.V1Init();
      //expect(instance).to.be();
    });

    it('should have the property tensorboard (base name: "tensorboard")', function() {
      // uncomment below and update the code to test the property tensorboard
      //var instance = new PolyaxonSdk.V1Init();
      //expect(instance).to.be();
    });

    it('should have the property lineageRef (base name: "lineageRef")', function() {
      // uncomment below and update the code to test the property lineageRef
      //var instance = new PolyaxonSdk.V1Init();
      //expect(instance).to.be();
    });

    it('should have the property artifactRef (base name: "artifactRef")', function() {
      // uncomment below and update the code to test the property artifactRef
      //var instance = new PolyaxonSdk.V1Init();
      //expect(instance).to.be();
    });

    it('should have the property modelRef (base name: "modelRef")', function() {
      // uncomment below and update the code to test the property modelRef
      //var instance = new PolyaxonSdk.V1Init();
      //expect(instance).to.be();
    });

    it('should have the property connection (base name: "connection")', function() {
      // uncomment below and update the code to test the property connection
      //var instance = new PolyaxonSdk.V1Init();
      //expect(instance).to.be();
    });

    it('should have the property path (base name: "path")', function() {
      // uncomment below and update the code to test the property path
      //var instance = new PolyaxonSdk.V1Init();
      //expect(instance).to.be();
    });

    it('should have the property container (base name: "container")', function() {
      // uncomment below and update the code to test the property container
      //var instance = new PolyaxonSdk.V1Init();
      //expect(instance).to.be();
    });

  });

}));
