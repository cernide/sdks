/*
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.IOException;
import org.openapitools.client.model.V1Dag;
import org.openapitools.client.model.V1DaskJob;
import org.openapitools.client.model.V1Job;
import org.openapitools.client.model.V1MPIJob;
import org.openapitools.client.model.V1MXJob;
import org.openapitools.client.model.V1PaddleJob;
import org.openapitools.client.model.V1PytorchJob;
import org.openapitools.client.model.V1RayJob;
import org.openapitools.client.model.V1Service;
import org.openapitools.client.model.V1TFJob;
import org.openapitools.client.model.V1XGBoostJob;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.google.gson.JsonArray;
import com.google.gson.JsonDeserializationContext;
import com.google.gson.JsonDeserializer;
import com.google.gson.JsonElement;
import com.google.gson.JsonObject;
import com.google.gson.JsonParseException;
import com.google.gson.TypeAdapterFactory;
import com.google.gson.reflect.TypeToken;

import java.lang.reflect.Type;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Map.Entry;
import java.util.Set;

import org.openapitools.client.JSON;

/**
 * V1RunSchema
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class V1RunSchema {
  public static final String SERIALIZED_NAME_JOB = "job";
  @SerializedName(SERIALIZED_NAME_JOB)
  private V1Job job;

  public static final String SERIALIZED_NAME_SERVICE = "service";
  @SerializedName(SERIALIZED_NAME_SERVICE)
  private V1Service service;

  public static final String SERIALIZED_NAME_DAG = "dag";
  @SerializedName(SERIALIZED_NAME_DAG)
  private V1Dag dag;

  public static final String SERIALIZED_NAME_TF_JOB = "tfJob";
  @SerializedName(SERIALIZED_NAME_TF_JOB)
  private V1TFJob tfJob;

  public static final String SERIALIZED_NAME_PYTORCH_JOB = "pytorchJob";
  @SerializedName(SERIALIZED_NAME_PYTORCH_JOB)
  private V1PytorchJob pytorchJob;

  public static final String SERIALIZED_NAME_MPI_JOB = "mpiJob";
  @SerializedName(SERIALIZED_NAME_MPI_JOB)
  private V1MPIJob mpiJob;

  public static final String SERIALIZED_NAME_MX_JOB = "mxJob";
  @SerializedName(SERIALIZED_NAME_MX_JOB)
  private V1MXJob mxJob;

  public static final String SERIALIZED_NAME_XGBOOST_JOB = "xgboostJob";
  @SerializedName(SERIALIZED_NAME_XGBOOST_JOB)
  private V1XGBoostJob xgboostJob;

  public static final String SERIALIZED_NAME_PADDLE_JOB = "paddleJob";
  @SerializedName(SERIALIZED_NAME_PADDLE_JOB)
  private V1PaddleJob paddleJob;

  public static final String SERIALIZED_NAME_DASK_JOB = "daskJob";
  @SerializedName(SERIALIZED_NAME_DASK_JOB)
  private V1DaskJob daskJob;

  public static final String SERIALIZED_NAME_RAY_JOB = "rayJob";
  @SerializedName(SERIALIZED_NAME_RAY_JOB)
  private V1RayJob rayJob;

  public V1RunSchema() {
  }

  public V1RunSchema job(V1Job job) {

    this.job = job;
    return this;
  }

   /**
   * Get job
   * @return job
  **/
  @javax.annotation.Nullable

  public V1Job getJob() {
    return job;
  }


  public void setJob(V1Job job) {
    this.job = job;
  }


  public V1RunSchema service(V1Service service) {

    this.service = service;
    return this;
  }

   /**
   * Get service
   * @return service
  **/
  @javax.annotation.Nullable

  public V1Service getService() {
    return service;
  }


  public void setService(V1Service service) {
    this.service = service;
  }


  public V1RunSchema dag(V1Dag dag) {

    this.dag = dag;
    return this;
  }

   /**
   * Get dag
   * @return dag
  **/
  @javax.annotation.Nullable

  public V1Dag getDag() {
    return dag;
  }


  public void setDag(V1Dag dag) {
    this.dag = dag;
  }


  public V1RunSchema tfJob(V1TFJob tfJob) {

    this.tfJob = tfJob;
    return this;
  }

   /**
   * Get tfJob
   * @return tfJob
  **/
  @javax.annotation.Nullable

  public V1TFJob getTfJob() {
    return tfJob;
  }


  public void setTfJob(V1TFJob tfJob) {
    this.tfJob = tfJob;
  }


  public V1RunSchema pytorchJob(V1PytorchJob pytorchJob) {

    this.pytorchJob = pytorchJob;
    return this;
  }

   /**
   * Get pytorchJob
   * @return pytorchJob
  **/
  @javax.annotation.Nullable

  public V1PytorchJob getPytorchJob() {
    return pytorchJob;
  }


  public void setPytorchJob(V1PytorchJob pytorchJob) {
    this.pytorchJob = pytorchJob;
  }


  public V1RunSchema mpiJob(V1MPIJob mpiJob) {

    this.mpiJob = mpiJob;
    return this;
  }

   /**
   * Get mpiJob
   * @return mpiJob
  **/
  @javax.annotation.Nullable

  public V1MPIJob getMpiJob() {
    return mpiJob;
  }


  public void setMpiJob(V1MPIJob mpiJob) {
    this.mpiJob = mpiJob;
  }


  public V1RunSchema mxJob(V1MXJob mxJob) {

    this.mxJob = mxJob;
    return this;
  }

   /**
   * Get mxJob
   * @return mxJob
  **/
  @javax.annotation.Nullable

  public V1MXJob getMxJob() {
    return mxJob;
  }


  public void setMxJob(V1MXJob mxJob) {
    this.mxJob = mxJob;
  }


  public V1RunSchema xgboostJob(V1XGBoostJob xgboostJob) {

    this.xgboostJob = xgboostJob;
    return this;
  }

   /**
   * Get xgboostJob
   * @return xgboostJob
  **/
  @javax.annotation.Nullable

  public V1XGBoostJob getXgboostJob() {
    return xgboostJob;
  }


  public void setXgboostJob(V1XGBoostJob xgboostJob) {
    this.xgboostJob = xgboostJob;
  }


  public V1RunSchema paddleJob(V1PaddleJob paddleJob) {

    this.paddleJob = paddleJob;
    return this;
  }

   /**
   * Get paddleJob
   * @return paddleJob
  **/
  @javax.annotation.Nullable

  public V1PaddleJob getPaddleJob() {
    return paddleJob;
  }


  public void setPaddleJob(V1PaddleJob paddleJob) {
    this.paddleJob = paddleJob;
  }


  public V1RunSchema daskJob(V1DaskJob daskJob) {

    this.daskJob = daskJob;
    return this;
  }

   /**
   * Get daskJob
   * @return daskJob
  **/
  @javax.annotation.Nullable

  public V1DaskJob getDaskJob() {
    return daskJob;
  }


  public void setDaskJob(V1DaskJob daskJob) {
    this.daskJob = daskJob;
  }


  public V1RunSchema rayJob(V1RayJob rayJob) {

    this.rayJob = rayJob;
    return this;
  }

   /**
   * Get rayJob
   * @return rayJob
  **/
  @javax.annotation.Nullable

  public V1RayJob getRayJob() {
    return rayJob;
  }


  public void setRayJob(V1RayJob rayJob) {
    this.rayJob = rayJob;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1RunSchema v1RunSchema = (V1RunSchema) o;
    return Objects.equals(this.job, v1RunSchema.job) &&
        Objects.equals(this.service, v1RunSchema.service) &&
        Objects.equals(this.dag, v1RunSchema.dag) &&
        Objects.equals(this.tfJob, v1RunSchema.tfJob) &&
        Objects.equals(this.pytorchJob, v1RunSchema.pytorchJob) &&
        Objects.equals(this.mpiJob, v1RunSchema.mpiJob) &&
        Objects.equals(this.mxJob, v1RunSchema.mxJob) &&
        Objects.equals(this.xgboostJob, v1RunSchema.xgboostJob) &&
        Objects.equals(this.paddleJob, v1RunSchema.paddleJob) &&
        Objects.equals(this.daskJob, v1RunSchema.daskJob) &&
        Objects.equals(this.rayJob, v1RunSchema.rayJob);
  }

  @Override
  public int hashCode() {
    return Objects.hash(job, service, dag, tfJob, pytorchJob, mpiJob, mxJob, xgboostJob, paddleJob, daskJob, rayJob);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1RunSchema {\n");
    sb.append("    job: ").append(toIndentedString(job)).append("\n");
    sb.append("    service: ").append(toIndentedString(service)).append("\n");
    sb.append("    dag: ").append(toIndentedString(dag)).append("\n");
    sb.append("    tfJob: ").append(toIndentedString(tfJob)).append("\n");
    sb.append("    pytorchJob: ").append(toIndentedString(pytorchJob)).append("\n");
    sb.append("    mpiJob: ").append(toIndentedString(mpiJob)).append("\n");
    sb.append("    mxJob: ").append(toIndentedString(mxJob)).append("\n");
    sb.append("    xgboostJob: ").append(toIndentedString(xgboostJob)).append("\n");
    sb.append("    paddleJob: ").append(toIndentedString(paddleJob)).append("\n");
    sb.append("    daskJob: ").append(toIndentedString(daskJob)).append("\n");
    sb.append("    rayJob: ").append(toIndentedString(rayJob)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }


  public static HashSet<String> openapiFields;
  public static HashSet<String> openapiRequiredFields;

  static {
    // a set of all properties/fields (JSON key names)
    openapiFields = new HashSet<String>();
    openapiFields.add("job");
    openapiFields.add("service");
    openapiFields.add("dag");
    openapiFields.add("tfJob");
    openapiFields.add("pytorchJob");
    openapiFields.add("mpiJob");
    openapiFields.add("mxJob");
    openapiFields.add("xgboostJob");
    openapiFields.add("paddleJob");
    openapiFields.add("daskJob");
    openapiFields.add("rayJob");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to V1RunSchema
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!V1RunSchema.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in V1RunSchema is not found in the empty JSON string", V1RunSchema.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!V1RunSchema.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `V1RunSchema` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }
      // validate the optional field `job`
      if (jsonObj.get("job") != null && !jsonObj.get("job").isJsonNull()) {
        V1Job.validateJsonObject(jsonObj.getAsJsonObject("job"));
      }
      // validate the optional field `service`
      if (jsonObj.get("service") != null && !jsonObj.get("service").isJsonNull()) {
        V1Service.validateJsonObject(jsonObj.getAsJsonObject("service"));
      }
      // validate the optional field `dag`
      if (jsonObj.get("dag") != null && !jsonObj.get("dag").isJsonNull()) {
        V1Dag.validateJsonObject(jsonObj.getAsJsonObject("dag"));
      }
      // validate the optional field `tfJob`
      if (jsonObj.get("tfJob") != null && !jsonObj.get("tfJob").isJsonNull()) {
        V1TFJob.validateJsonObject(jsonObj.getAsJsonObject("tfJob"));
      }
      // validate the optional field `pytorchJob`
      if (jsonObj.get("pytorchJob") != null && !jsonObj.get("pytorchJob").isJsonNull()) {
        V1PytorchJob.validateJsonObject(jsonObj.getAsJsonObject("pytorchJob"));
      }
      // validate the optional field `mpiJob`
      if (jsonObj.get("mpiJob") != null && !jsonObj.get("mpiJob").isJsonNull()) {
        V1MPIJob.validateJsonObject(jsonObj.getAsJsonObject("mpiJob"));
      }
      // validate the optional field `mxJob`
      if (jsonObj.get("mxJob") != null && !jsonObj.get("mxJob").isJsonNull()) {
        V1MXJob.validateJsonObject(jsonObj.getAsJsonObject("mxJob"));
      }
      // validate the optional field `xgboostJob`
      if (jsonObj.get("xgboostJob") != null && !jsonObj.get("xgboostJob").isJsonNull()) {
        V1XGBoostJob.validateJsonObject(jsonObj.getAsJsonObject("xgboostJob"));
      }
      // validate the optional field `paddleJob`
      if (jsonObj.get("paddleJob") != null && !jsonObj.get("paddleJob").isJsonNull()) {
        V1PaddleJob.validateJsonObject(jsonObj.getAsJsonObject("paddleJob"));
      }
      // validate the optional field `daskJob`
      if (jsonObj.get("daskJob") != null && !jsonObj.get("daskJob").isJsonNull()) {
        V1DaskJob.validateJsonObject(jsonObj.getAsJsonObject("daskJob"));
      }
      // validate the optional field `rayJob`
      if (jsonObj.get("rayJob") != null && !jsonObj.get("rayJob").isJsonNull()) {
        V1RayJob.validateJsonObject(jsonObj.getAsJsonObject("rayJob"));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!V1RunSchema.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'V1RunSchema' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<V1RunSchema> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(V1RunSchema.class));

       return (TypeAdapter<T>) new TypeAdapter<V1RunSchema>() {
           @Override
           public void write(JsonWriter out, V1RunSchema value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public V1RunSchema read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of V1RunSchema given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of V1RunSchema
  * @throws IOException if the JSON string is invalid with respect to V1RunSchema
  */
  public static V1RunSchema fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, V1RunSchema.class);
  }

 /**
  * Convert an instance of V1RunSchema to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

