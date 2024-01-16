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
    instance = new PolyaxonSdk.ProjectDashboardsV1Api();
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

  describe('ProjectDashboardsV1Api', function() {
    describe('createProjectDashboard', function() {
      it('should call createProjectDashboard successfully', function(done) {
        //uncomment below and update the code to test createProjectDashboard
        //instance.createProjectDashboard(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('deleteProjectDashboard', function() {
      it('should call deleteProjectDashboard successfully', function(done) {
        //uncomment below and update the code to test deleteProjectDashboard
        //instance.deleteProjectDashboard(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getProjectDashboard', function() {
      it('should call getProjectDashboard successfully', function(done) {
        //uncomment below and update the code to test getProjectDashboard
        //instance.getProjectDashboard(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listProjectDashboardNames', function() {
      it('should call listProjectDashboardNames successfully', function(done) {
        //uncomment below and update the code to test listProjectDashboardNames
        //instance.listProjectDashboardNames(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listProjectDashboards', function() {
      it('should call listProjectDashboards successfully', function(done) {
        //uncomment below and update the code to test listProjectDashboards
        //instance.listProjectDashboards(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('patchProjectDashboard', function() {
      it('should call patchProjectDashboard successfully', function(done) {
        //uncomment below and update the code to test patchProjectDashboard
        //instance.patchProjectDashboard(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('promoteProjectDashboard', function() {
      it('should call promoteProjectDashboard successfully', function(done) {
        //uncomment below and update the code to test promoteProjectDashboard
        //instance.promoteProjectDashboard(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateProjectDashboard', function() {
      it('should call updateProjectDashboard successfully', function(done) {
        //uncomment below and update the code to test updateProjectDashboard
        //instance.updateProjectDashboard(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
  });

}));
