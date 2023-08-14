/*
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0-rc36
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
 * V1AnalyticsSpec
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class V1AnalyticsSpec {
  public static final String SERIALIZED_NAME_VIEW = "view";
  @SerializedName(SERIALIZED_NAME_VIEW)
  private String view;

  public static final String SERIALIZED_NAME_TRUNC = "trunc";
  @SerializedName(SERIALIZED_NAME_TRUNC)
  private String trunc;

  public static final String SERIALIZED_NAME_GROUPBY = "groupby";
  @SerializedName(SERIALIZED_NAME_GROUPBY)
  private String groupby;

  public static final String SERIALIZED_NAME_FREQUENCY = "frequency";
  @SerializedName(SERIALIZED_NAME_FREQUENCY)
  private String frequency;

  public V1AnalyticsSpec() {
  }

  public V1AnalyticsSpec view(String view) {

    this.view = view;
    return this;
  }

   /**
   * Get view
   * @return view
  **/
  @javax.annotation.Nullable

  public String getView() {
    return view;
  }


  public void setView(String view) {
    this.view = view;
  }


  public V1AnalyticsSpec trunc(String trunc) {

    this.trunc = trunc;
    return this;
  }

   /**
   * Get trunc
   * @return trunc
  **/
  @javax.annotation.Nullable

  public String getTrunc() {
    return trunc;
  }


  public void setTrunc(String trunc) {
    this.trunc = trunc;
  }


  public V1AnalyticsSpec groupby(String groupby) {

    this.groupby = groupby;
    return this;
  }

   /**
   * Get groupby
   * @return groupby
  **/
  @javax.annotation.Nullable

  public String getGroupby() {
    return groupby;
  }


  public void setGroupby(String groupby) {
    this.groupby = groupby;
  }


  public V1AnalyticsSpec frequency(String frequency) {

    this.frequency = frequency;
    return this;
  }

   /**
   * Get frequency
   * @return frequency
  **/
  @javax.annotation.Nullable

  public String getFrequency() {
    return frequency;
  }


  public void setFrequency(String frequency) {
    this.frequency = frequency;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1AnalyticsSpec v1AnalyticsSpec = (V1AnalyticsSpec) o;
    return Objects.equals(this.view, v1AnalyticsSpec.view) &&
        Objects.equals(this.trunc, v1AnalyticsSpec.trunc) &&
        Objects.equals(this.groupby, v1AnalyticsSpec.groupby) &&
        Objects.equals(this.frequency, v1AnalyticsSpec.frequency);
  }

  @Override
  public int hashCode() {
    return Objects.hash(view, trunc, groupby, frequency);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1AnalyticsSpec {\n");
    sb.append("    view: ").append(toIndentedString(view)).append("\n");
    sb.append("    trunc: ").append(toIndentedString(trunc)).append("\n");
    sb.append("    groupby: ").append(toIndentedString(groupby)).append("\n");
    sb.append("    frequency: ").append(toIndentedString(frequency)).append("\n");
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
    openapiFields.add("view");
    openapiFields.add("trunc");
    openapiFields.add("groupby");
    openapiFields.add("frequency");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to V1AnalyticsSpec
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!V1AnalyticsSpec.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in V1AnalyticsSpec is not found in the empty JSON string", V1AnalyticsSpec.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!V1AnalyticsSpec.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `V1AnalyticsSpec` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }
      if ((jsonObj.get("view") != null && !jsonObj.get("view").isJsonNull()) && !jsonObj.get("view").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `view` to be a primitive type in the JSON string but got `%s`", jsonObj.get("view").toString()));
      }
      if ((jsonObj.get("trunc") != null && !jsonObj.get("trunc").isJsonNull()) && !jsonObj.get("trunc").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `trunc` to be a primitive type in the JSON string but got `%s`", jsonObj.get("trunc").toString()));
      }
      if ((jsonObj.get("groupby") != null && !jsonObj.get("groupby").isJsonNull()) && !jsonObj.get("groupby").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `groupby` to be a primitive type in the JSON string but got `%s`", jsonObj.get("groupby").toString()));
      }
      if ((jsonObj.get("frequency") != null && !jsonObj.get("frequency").isJsonNull()) && !jsonObj.get("frequency").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `frequency` to be a primitive type in the JSON string but got `%s`", jsonObj.get("frequency").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!V1AnalyticsSpec.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'V1AnalyticsSpec' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<V1AnalyticsSpec> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(V1AnalyticsSpec.class));

       return (TypeAdapter<T>) new TypeAdapter<V1AnalyticsSpec>() {
           @Override
           public void write(JsonWriter out, V1AnalyticsSpec value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public V1AnalyticsSpec read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of V1AnalyticsSpec given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of V1AnalyticsSpec
  * @throws IOException if the JSON string is invalid with respect to V1AnalyticsSpec
  */
  public static V1AnalyticsSpec fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, V1AnalyticsSpec.class);
  }

 /**
  * Convert an instance of V1AnalyticsSpec to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

