/*
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.4.0
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
import org.openapitools.client.model.V1DaskReplica;

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
 * V1DaskJob
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class V1DaskJob {
  public static final String SERIALIZED_NAME_KIND = "kind";
  @SerializedName(SERIALIZED_NAME_KIND)
  private String kind = "daskjob";

  public static final String SERIALIZED_NAME_JOB = "job";
  @SerializedName(SERIALIZED_NAME_JOB)
  private V1DaskReplica job;

  public static final String SERIALIZED_NAME_WORKER = "worker";
  @SerializedName(SERIALIZED_NAME_WORKER)
  private V1DaskReplica worker;

  public static final String SERIALIZED_NAME_SCHEDULER = "scheduler";
  @SerializedName(SERIALIZED_NAME_SCHEDULER)
  private V1DaskReplica scheduler;

  public V1DaskJob() {
  }

  public V1DaskJob kind(String kind) {

    this.kind = kind;
    return this;
  }

   /**
   * Get kind
   * @return kind
  **/
  @javax.annotation.Nullable

  public String getKind() {
    return kind;
  }


  public void setKind(String kind) {
    this.kind = kind;
  }


  public V1DaskJob job(V1DaskReplica job) {

    this.job = job;
    return this;
  }

   /**
   * Get job
   * @return job
  **/
  @javax.annotation.Nullable

  public V1DaskReplica getJob() {
    return job;
  }


  public void setJob(V1DaskReplica job) {
    this.job = job;
  }


  public V1DaskJob worker(V1DaskReplica worker) {

    this.worker = worker;
    return this;
  }

   /**
   * Get worker
   * @return worker
  **/
  @javax.annotation.Nullable

  public V1DaskReplica getWorker() {
    return worker;
  }


  public void setWorker(V1DaskReplica worker) {
    this.worker = worker;
  }


  public V1DaskJob scheduler(V1DaskReplica scheduler) {

    this.scheduler = scheduler;
    return this;
  }

   /**
   * Get scheduler
   * @return scheduler
  **/
  @javax.annotation.Nullable

  public V1DaskReplica getScheduler() {
    return scheduler;
  }


  public void setScheduler(V1DaskReplica scheduler) {
    this.scheduler = scheduler;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1DaskJob v1DaskJob = (V1DaskJob) o;
    return Objects.equals(this.kind, v1DaskJob.kind) &&
        Objects.equals(this.job, v1DaskJob.job) &&
        Objects.equals(this.worker, v1DaskJob.worker) &&
        Objects.equals(this.scheduler, v1DaskJob.scheduler);
  }

  @Override
  public int hashCode() {
    return Objects.hash(kind, job, worker, scheduler);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1DaskJob {\n");
    sb.append("    kind: ").append(toIndentedString(kind)).append("\n");
    sb.append("    job: ").append(toIndentedString(job)).append("\n");
    sb.append("    worker: ").append(toIndentedString(worker)).append("\n");
    sb.append("    scheduler: ").append(toIndentedString(scheduler)).append("\n");
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
    openapiFields.add("kind");
    openapiFields.add("job");
    openapiFields.add("worker");
    openapiFields.add("scheduler");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to V1DaskJob
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!V1DaskJob.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in V1DaskJob is not found in the empty JSON string", V1DaskJob.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!V1DaskJob.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `V1DaskJob` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }
      if ((jsonObj.get("kind") != null && !jsonObj.get("kind").isJsonNull()) && !jsonObj.get("kind").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `kind` to be a primitive type in the JSON string but got `%s`", jsonObj.get("kind").toString()));
      }
      // validate the optional field `job`
      if (jsonObj.get("job") != null && !jsonObj.get("job").isJsonNull()) {
        V1DaskReplica.validateJsonObject(jsonObj.getAsJsonObject("job"));
      }
      // validate the optional field `worker`
      if (jsonObj.get("worker") != null && !jsonObj.get("worker").isJsonNull()) {
        V1DaskReplica.validateJsonObject(jsonObj.getAsJsonObject("worker"));
      }
      // validate the optional field `scheduler`
      if (jsonObj.get("scheduler") != null && !jsonObj.get("scheduler").isJsonNull()) {
        V1DaskReplica.validateJsonObject(jsonObj.getAsJsonObject("scheduler"));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!V1DaskJob.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'V1DaskJob' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<V1DaskJob> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(V1DaskJob.class));

       return (TypeAdapter<T>) new TypeAdapter<V1DaskJob>() {
           @Override
           public void write(JsonWriter out, V1DaskJob value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public V1DaskJob read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of V1DaskJob given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of V1DaskJob
  * @throws IOException if the JSON string is invalid with respect to V1DaskJob
  */
  public static V1DaskJob fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, V1DaskJob.class);
  }

 /**
  * Convert an instance of V1DaskJob to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

