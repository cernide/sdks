/*
 * Polyaxon SDKs and REST API specification.
 *    
 *
 * The version of the OpenAPI document: 2.1.6-rc0
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
import java.time.OffsetDateTime;

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
 * V1IntervalSchedule
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class V1IntervalSchedule {
  public static final String SERIALIZED_NAME_KIND = "kind";
  @SerializedName(SERIALIZED_NAME_KIND)
  private String kind = "interval";

  public static final String SERIALIZED_NAME_START_AT = "startAt";
  @SerializedName(SERIALIZED_NAME_START_AT)
  private OffsetDateTime startAt;

  public static final String SERIALIZED_NAME_END_AT = "endAt";
  @SerializedName(SERIALIZED_NAME_END_AT)
  private OffsetDateTime endAt;

  public static final String SERIALIZED_NAME_MAX_RUNS = "maxRuns";
  @SerializedName(SERIALIZED_NAME_MAX_RUNS)
  private Integer maxRuns;

  public static final String SERIALIZED_NAME_FREQUENCY = "frequency";
  @SerializedName(SERIALIZED_NAME_FREQUENCY)
  private Integer frequency;

  public static final String SERIALIZED_NAME_DEPENDS_ON_PAST = "dependsOnPast";
  @SerializedName(SERIALIZED_NAME_DEPENDS_ON_PAST)
  private Boolean dependsOnPast;

  public V1IntervalSchedule() {
  }

  public V1IntervalSchedule kind(String kind) {
    
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


  public V1IntervalSchedule startAt(OffsetDateTime startAt) {
    
    this.startAt = startAt;
    return this;
  }

   /**
   * Get startAt
   * @return startAt
  **/
  @javax.annotation.Nullable

  public OffsetDateTime getStartAt() {
    return startAt;
  }


  public void setStartAt(OffsetDateTime startAt) {
    this.startAt = startAt;
  }


  public V1IntervalSchedule endAt(OffsetDateTime endAt) {
    
    this.endAt = endAt;
    return this;
  }

   /**
   * Get endAt
   * @return endAt
  **/
  @javax.annotation.Nullable

  public OffsetDateTime getEndAt() {
    return endAt;
  }


  public void setEndAt(OffsetDateTime endAt) {
    this.endAt = endAt;
  }


  public V1IntervalSchedule maxRuns(Integer maxRuns) {
    
    this.maxRuns = maxRuns;
    return this;
  }

   /**
   * Get maxRuns
   * @return maxRuns
  **/
  @javax.annotation.Nullable

  public Integer getMaxRuns() {
    return maxRuns;
  }


  public void setMaxRuns(Integer maxRuns) {
    this.maxRuns = maxRuns;
  }


  public V1IntervalSchedule frequency(Integer frequency) {
    
    this.frequency = frequency;
    return this;
  }

   /**
   * Get frequency
   * @return frequency
  **/
  @javax.annotation.Nullable

  public Integer getFrequency() {
    return frequency;
  }


  public void setFrequency(Integer frequency) {
    this.frequency = frequency;
  }


  public V1IntervalSchedule dependsOnPast(Boolean dependsOnPast) {
    
    this.dependsOnPast = dependsOnPast;
    return this;
  }

   /**
   * Get dependsOnPast
   * @return dependsOnPast
  **/
  @javax.annotation.Nullable

  public Boolean getDependsOnPast() {
    return dependsOnPast;
  }


  public void setDependsOnPast(Boolean dependsOnPast) {
    this.dependsOnPast = dependsOnPast;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1IntervalSchedule v1IntervalSchedule = (V1IntervalSchedule) o;
    return Objects.equals(this.kind, v1IntervalSchedule.kind) &&
        Objects.equals(this.startAt, v1IntervalSchedule.startAt) &&
        Objects.equals(this.endAt, v1IntervalSchedule.endAt) &&
        Objects.equals(this.maxRuns, v1IntervalSchedule.maxRuns) &&
        Objects.equals(this.frequency, v1IntervalSchedule.frequency) &&
        Objects.equals(this.dependsOnPast, v1IntervalSchedule.dependsOnPast);
  }

  @Override
  public int hashCode() {
    return Objects.hash(kind, startAt, endAt, maxRuns, frequency, dependsOnPast);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1IntervalSchedule {\n");
    sb.append("    kind: ").append(toIndentedString(kind)).append("\n");
    sb.append("    startAt: ").append(toIndentedString(startAt)).append("\n");
    sb.append("    endAt: ").append(toIndentedString(endAt)).append("\n");
    sb.append("    maxRuns: ").append(toIndentedString(maxRuns)).append("\n");
    sb.append("    frequency: ").append(toIndentedString(frequency)).append("\n");
    sb.append("    dependsOnPast: ").append(toIndentedString(dependsOnPast)).append("\n");
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
    openapiFields.add("startAt");
    openapiFields.add("endAt");
    openapiFields.add("maxRuns");
    openapiFields.add("frequency");
    openapiFields.add("dependsOnPast");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to V1IntervalSchedule
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!V1IntervalSchedule.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in V1IntervalSchedule is not found in the empty JSON string", V1IntervalSchedule.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!V1IntervalSchedule.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `V1IntervalSchedule` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }
      if ((jsonObj.get("kind") != null && !jsonObj.get("kind").isJsonNull()) && !jsonObj.get("kind").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `kind` to be a primitive type in the JSON string but got `%s`", jsonObj.get("kind").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!V1IntervalSchedule.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'V1IntervalSchedule' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<V1IntervalSchedule> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(V1IntervalSchedule.class));

       return (TypeAdapter<T>) new TypeAdapter<V1IntervalSchedule>() {
           @Override
           public void write(JsonWriter out, V1IntervalSchedule value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public V1IntervalSchedule read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of V1IntervalSchedule given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of V1IntervalSchedule
  * @throws IOException if the JSON string is invalid with respect to V1IntervalSchedule
  */
  public static V1IntervalSchedule fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, V1IntervalSchedule.class);
  }

 /**
  * Convert an instance of V1IntervalSchedule to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

