/*
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0-rc33
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
import org.openapitools.client.model.V1CleanPodPolicy;
import org.openapitools.client.model.V1KFReplica;
import org.openapitools.client.model.V1SchedulingPolicy;

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
 * V1PaddleJob
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class V1PaddleJob {
  public static final String SERIALIZED_NAME_KIND = "kind";
  @SerializedName(SERIALIZED_NAME_KIND)
  private String kind = "paddlejob";

  public static final String SERIALIZED_NAME_CLEAN_POD_POLICY = "cleanPodPolicy";
  @SerializedName(SERIALIZED_NAME_CLEAN_POD_POLICY)
  private V1CleanPodPolicy cleanPodPolicy = V1CleanPodPolicy.ALL;

  public static final String SERIALIZED_NAME_SCHEDULING_POLICY = "schedulingPolicy";
  @SerializedName(SERIALIZED_NAME_SCHEDULING_POLICY)
  private V1SchedulingPolicy schedulingPolicy;

  public static final String SERIALIZED_NAME_MASTER = "master";
  @SerializedName(SERIALIZED_NAME_MASTER)
  private V1KFReplica master;

  public static final String SERIALIZED_NAME_WORKER = "worker";
  @SerializedName(SERIALIZED_NAME_WORKER)
  private V1KFReplica worker;

  public V1PaddleJob() {
  }

  public V1PaddleJob kind(String kind) {

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


  public V1PaddleJob cleanPodPolicy(V1CleanPodPolicy cleanPodPolicy) {

    this.cleanPodPolicy = cleanPodPolicy;
    return this;
  }

   /**
   * Get cleanPodPolicy
   * @return cleanPodPolicy
  **/
  @javax.annotation.Nullable

  public V1CleanPodPolicy getCleanPodPolicy() {
    return cleanPodPolicy;
  }


  public void setCleanPodPolicy(V1CleanPodPolicy cleanPodPolicy) {
    this.cleanPodPolicy = cleanPodPolicy;
  }


  public V1PaddleJob schedulingPolicy(V1SchedulingPolicy schedulingPolicy) {

    this.schedulingPolicy = schedulingPolicy;
    return this;
  }

   /**
   * Get schedulingPolicy
   * @return schedulingPolicy
  **/
  @javax.annotation.Nullable

  public V1SchedulingPolicy getSchedulingPolicy() {
    return schedulingPolicy;
  }


  public void setSchedulingPolicy(V1SchedulingPolicy schedulingPolicy) {
    this.schedulingPolicy = schedulingPolicy;
  }


  public V1PaddleJob master(V1KFReplica master) {

    this.master = master;
    return this;
  }

   /**
   * Get master
   * @return master
  **/
  @javax.annotation.Nullable

  public V1KFReplica getMaster() {
    return master;
  }


  public void setMaster(V1KFReplica master) {
    this.master = master;
  }


  public V1PaddleJob worker(V1KFReplica worker) {

    this.worker = worker;
    return this;
  }

   /**
   * Get worker
   * @return worker
  **/
  @javax.annotation.Nullable

  public V1KFReplica getWorker() {
    return worker;
  }


  public void setWorker(V1KFReplica worker) {
    this.worker = worker;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1PaddleJob v1PaddleJob = (V1PaddleJob) o;
    return Objects.equals(this.kind, v1PaddleJob.kind) &&
        Objects.equals(this.cleanPodPolicy, v1PaddleJob.cleanPodPolicy) &&
        Objects.equals(this.schedulingPolicy, v1PaddleJob.schedulingPolicy) &&
        Objects.equals(this.master, v1PaddleJob.master) &&
        Objects.equals(this.worker, v1PaddleJob.worker);
  }

  @Override
  public int hashCode() {
    return Objects.hash(kind, cleanPodPolicy, schedulingPolicy, master, worker);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1PaddleJob {\n");
    sb.append("    kind: ").append(toIndentedString(kind)).append("\n");
    sb.append("    cleanPodPolicy: ").append(toIndentedString(cleanPodPolicy)).append("\n");
    sb.append("    schedulingPolicy: ").append(toIndentedString(schedulingPolicy)).append("\n");
    sb.append("    master: ").append(toIndentedString(master)).append("\n");
    sb.append("    worker: ").append(toIndentedString(worker)).append("\n");
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
    openapiFields.add("cleanPodPolicy");
    openapiFields.add("schedulingPolicy");
    openapiFields.add("master");
    openapiFields.add("worker");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to V1PaddleJob
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!V1PaddleJob.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in V1PaddleJob is not found in the empty JSON string", V1PaddleJob.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!V1PaddleJob.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `V1PaddleJob` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }
      if ((jsonObj.get("kind") != null && !jsonObj.get("kind").isJsonNull()) && !jsonObj.get("kind").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `kind` to be a primitive type in the JSON string but got `%s`", jsonObj.get("kind").toString()));
      }
      // validate the optional field `schedulingPolicy`
      if (jsonObj.get("schedulingPolicy") != null && !jsonObj.get("schedulingPolicy").isJsonNull()) {
        V1SchedulingPolicy.validateJsonObject(jsonObj.getAsJsonObject("schedulingPolicy"));
      }
      // validate the optional field `master`
      if (jsonObj.get("master") != null && !jsonObj.get("master").isJsonNull()) {
        V1KFReplica.validateJsonObject(jsonObj.getAsJsonObject("master"));
      }
      // validate the optional field `worker`
      if (jsonObj.get("worker") != null && !jsonObj.get("worker").isJsonNull()) {
        V1KFReplica.validateJsonObject(jsonObj.getAsJsonObject("worker"));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!V1PaddleJob.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'V1PaddleJob' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<V1PaddleJob> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(V1PaddleJob.class));

       return (TypeAdapter<T>) new TypeAdapter<V1PaddleJob>() {
           @Override
           public void write(JsonWriter out, V1PaddleJob value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public V1PaddleJob read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of V1PaddleJob given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of V1PaddleJob
  * @throws IOException if the JSON string is invalid with respect to V1PaddleJob
  */
  public static V1PaddleJob fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, V1PaddleJob.class);
  }

 /**
  * Convert an instance of V1PaddleJob to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

