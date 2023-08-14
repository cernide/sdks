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
    instance = new PolyaxonSdk.V1DockerfileType();
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

  describe('V1DockerfileType', function() {
    it('should create an instance of V1DockerfileType', function() {
      // uncomment below and update the code to test V1DockerfileType
      //var instance = new PolyaxonSdk.V1DockerfileType();
      //expect(instance).to.be.a(PolyaxonSdk.V1DockerfileType);
    });

    it('should have the property image (base name: "image")', function() {
      // uncomment below and update the code to test the property image
      //var instance = new PolyaxonSdk.V1DockerfileType();
      //expect(instance).to.be();
    });

    it('should have the property env (base name: "env")', function() {
      // uncomment below and update the code to test the property env
      //var instance = new PolyaxonSdk.V1DockerfileType();
      //expect(instance).to.be();
    });

    it('should have the property path (base name: "path")', function() {
      // uncomment below and update the code to test the property path
      //var instance = new PolyaxonSdk.V1DockerfileType();
      //expect(instance).to.be();
    });

    it('should have the property copy (base name: "copy")', function() {
      // uncomment below and update the code to test the property copy
      //var instance = new PolyaxonSdk.V1DockerfileType();
      //expect(instance).to.be();
    });

    it('should have the property post_run_copy (base name: "post_run_copy")', function() {
      // uncomment below and update the code to test the property post_run_copy
      //var instance = new PolyaxonSdk.V1DockerfileType();
      //expect(instance).to.be();
    });

    it('should have the property run (base name: "run")', function() {
      // uncomment below and update the code to test the property run
      //var instance = new PolyaxonSdk.V1DockerfileType();
      //expect(instance).to.be();
    });

    it('should have the property langEnv (base name: "langEnv")', function() {
      // uncomment below and update the code to test the property langEnv
      //var instance = new PolyaxonSdk.V1DockerfileType();
      //expect(instance).to.be();
    });

    it('should have the property uid (base name: "uid")', function() {
      // uncomment below and update the code to test the property uid
      //var instance = new PolyaxonSdk.V1DockerfileType();
      //expect(instance).to.be();
    });

    it('should have the property gid (base name: "gid")', function() {
      // uncomment below and update the code to test the property gid
      //var instance = new PolyaxonSdk.V1DockerfileType();
      //expect(instance).to.be();
    });

    it('should have the property username (base name: "username")', function() {
      // uncomment below and update the code to test the property username
      //var instance = new PolyaxonSdk.V1DockerfileType();
      //expect(instance).to.be();
    });

    it('should have the property filename (base name: "filename")', function() {
      // uncomment below and update the code to test the property filename
      //var instance = new PolyaxonSdk.V1DockerfileType();
      //expect(instance).to.be();
    });

    it('should have the property workdir (base name: "workdir")', function() {
      // uncomment below and update the code to test the property workdir
      //var instance = new PolyaxonSdk.V1DockerfileType();
      //expect(instance).to.be();
    });

    it('should have the property workdirPath (base name: "workdirPath")', function() {
      // uncomment below and update the code to test the property workdirPath
      //var instance = new PolyaxonSdk.V1DockerfileType();
      //expect(instance).to.be();
    });

    it('should have the property shell (base name: "shell")', function() {
      // uncomment below and update the code to test the property shell
      //var instance = new PolyaxonSdk.V1DockerfileType();
      //expect(instance).to.be();
    });

  });

}));
