// Copyright 2018-2021 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/**
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * The version of the OpenAPI document: 1.13.0-rc3
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
    instance = new PolyaxonSdk.V1Plugins();
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

  describe('V1Plugins', function() {
    it('should create an instance of V1Plugins', function() {
      // uncomment below and update the code to test V1Plugins
      //var instance = new PolyaxonSdk.V1Plugins();
      //expect(instance).to.be.a(PolyaxonSdk.V1Plugins);
    });

    it('should have the property auth (base name: "auth")', function() {
      // uncomment below and update the code to test the property auth
      //var instance = new PolyaxonSdk.V1Plugins();
      //expect(instance).to.be();
    });

    it('should have the property docker (base name: "docker")', function() {
      // uncomment below and update the code to test the property docker
      //var instance = new PolyaxonSdk.V1Plugins();
      //expect(instance).to.be();
    });

    it('should have the property shm (base name: "shm")', function() {
      // uncomment below and update the code to test the property shm
      //var instance = new PolyaxonSdk.V1Plugins();
      //expect(instance).to.be();
    });

    it('should have the property mountArtifactsStore (base name: "mountArtifactsStore")', function() {
      // uncomment below and update the code to test the property mountArtifactsStore
      //var instance = new PolyaxonSdk.V1Plugins();
      //expect(instance).to.be();
    });

    it('should have the property collectArtifacts (base name: "collectArtifacts")', function() {
      // uncomment below and update the code to test the property collectArtifacts
      //var instance = new PolyaxonSdk.V1Plugins();
      //expect(instance).to.be();
    });

    it('should have the property collectLogs (base name: "collectLogs")', function() {
      // uncomment below and update the code to test the property collectLogs
      //var instance = new PolyaxonSdk.V1Plugins();
      //expect(instance).to.be();
    });

    it('should have the property collectResources (base name: "collectResources")', function() {
      // uncomment below and update the code to test the property collectResources
      //var instance = new PolyaxonSdk.V1Plugins();
      //expect(instance).to.be();
    });

    it('should have the property syncStatuses (base name: "syncStatuses")', function() {
      // uncomment below and update the code to test the property syncStatuses
      //var instance = new PolyaxonSdk.V1Plugins();
      //expect(instance).to.be();
    });

    it('should have the property autoResume (base name: "autoResume")', function() {
      // uncomment below and update the code to test the property autoResume
      //var instance = new PolyaxonSdk.V1Plugins();
      //expect(instance).to.be();
    });

    it('should have the property logLevel (base name: "logLevel")', function() {
      // uncomment below and update the code to test the property logLevel
      //var instance = new PolyaxonSdk.V1Plugins();
      //expect(instance).to.be();
    });

    it('should have the property externalHost (base name: "externalHost")', function() {
      // uncomment below and update the code to test the property externalHost
      //var instance = new PolyaxonSdk.V1Plugins();
      //expect(instance).to.be();
    });

    it('should have the property sidecar (base name: "sidecar")', function() {
      // uncomment below and update the code to test the property sidecar
      //var instance = new PolyaxonSdk.V1Plugins();
      //expect(instance).to.be();
    });

    it('should have the property notifications (base name: "notifications")', function() {
      // uncomment below and update the code to test the property notifications
      //var instance = new PolyaxonSdk.V1Plugins();
      //expect(instance).to.be();
    });

  });

}));