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
import org.openapitools.client.model.V1DagRef;
import org.openapitools.client.model.V1HubRef;
import org.openapitools.client.model.V1PathRef;
import org.openapitools.client.model.V1UrlRef;

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
 * V1Reference
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class V1Reference {
  public static final String SERIALIZED_NAME_HUB_REF = "hubRef";
  @SerializedName(SERIALIZED_NAME_HUB_REF)
  private V1HubRef hubRef;

  public static final String SERIALIZED_NAME_DAG_REF = "dagRef";
  @SerializedName(SERIALIZED_NAME_DAG_REF)
  private V1DagRef dagRef;

  public static final String SERIALIZED_NAME_URL_REF = "urlRef";
  @SerializedName(SERIALIZED_NAME_URL_REF)
  private V1UrlRef urlRef;

  public static final String SERIALIZED_NAME_PATH_REF = "pathRef";
  @SerializedName(SERIALIZED_NAME_PATH_REF)
  private V1PathRef pathRef;

  public V1Reference() {
  }

  public V1Reference hubRef(V1HubRef hubRef) {

    this.hubRef = hubRef;
    return this;
  }

   /**
   * Get hubRef
   * @return hubRef
  **/
  @javax.annotation.Nullable

  public V1HubRef getHubRef() {
    return hubRef;
  }


  public void setHubRef(V1HubRef hubRef) {
    this.hubRef = hubRef;
  }


  public V1Reference dagRef(V1DagRef dagRef) {

    this.dagRef = dagRef;
    return this;
  }

   /**
   * Get dagRef
   * @return dagRef
  **/
  @javax.annotation.Nullable

  public V1DagRef getDagRef() {
    return dagRef;
  }


  public void setDagRef(V1DagRef dagRef) {
    this.dagRef = dagRef;
  }


  public V1Reference urlRef(V1UrlRef urlRef) {

    this.urlRef = urlRef;
    return this;
  }

   /**
   * Get urlRef
   * @return urlRef
  **/
  @javax.annotation.Nullable

  public V1UrlRef getUrlRef() {
    return urlRef;
  }


  public void setUrlRef(V1UrlRef urlRef) {
    this.urlRef = urlRef;
  }


  public V1Reference pathRef(V1PathRef pathRef) {

    this.pathRef = pathRef;
    return this;
  }

   /**
   * Get pathRef
   * @return pathRef
  **/
  @javax.annotation.Nullable

  public V1PathRef getPathRef() {
    return pathRef;
  }


  public void setPathRef(V1PathRef pathRef) {
    this.pathRef = pathRef;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1Reference v1Reference = (V1Reference) o;
    return Objects.equals(this.hubRef, v1Reference.hubRef) &&
        Objects.equals(this.dagRef, v1Reference.dagRef) &&
        Objects.equals(this.urlRef, v1Reference.urlRef) &&
        Objects.equals(this.pathRef, v1Reference.pathRef);
  }

  @Override
  public int hashCode() {
    return Objects.hash(hubRef, dagRef, urlRef, pathRef);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1Reference {\n");
    sb.append("    hubRef: ").append(toIndentedString(hubRef)).append("\n");
    sb.append("    dagRef: ").append(toIndentedString(dagRef)).append("\n");
    sb.append("    urlRef: ").append(toIndentedString(urlRef)).append("\n");
    sb.append("    pathRef: ").append(toIndentedString(pathRef)).append("\n");
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
    openapiFields.add("hubRef");
    openapiFields.add("dagRef");
    openapiFields.add("urlRef");
    openapiFields.add("pathRef");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to V1Reference
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!V1Reference.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in V1Reference is not found in the empty JSON string", V1Reference.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!V1Reference.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `V1Reference` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }
      // validate the optional field `hubRef`
      if (jsonObj.get("hubRef") != null && !jsonObj.get("hubRef").isJsonNull()) {
        V1HubRef.validateJsonObject(jsonObj.getAsJsonObject("hubRef"));
      }
      // validate the optional field `dagRef`
      if (jsonObj.get("dagRef") != null && !jsonObj.get("dagRef").isJsonNull()) {
        V1DagRef.validateJsonObject(jsonObj.getAsJsonObject("dagRef"));
      }
      // validate the optional field `urlRef`
      if (jsonObj.get("urlRef") != null && !jsonObj.get("urlRef").isJsonNull()) {
        V1UrlRef.validateJsonObject(jsonObj.getAsJsonObject("urlRef"));
      }
      // validate the optional field `pathRef`
      if (jsonObj.get("pathRef") != null && !jsonObj.get("pathRef").isJsonNull()) {
        V1PathRef.validateJsonObject(jsonObj.getAsJsonObject("pathRef"));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!V1Reference.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'V1Reference' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<V1Reference> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(V1Reference.class));

       return (TypeAdapter<T>) new TypeAdapter<V1Reference>() {
           @Override
           public void write(JsonWriter out, V1Reference value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public V1Reference read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of V1Reference given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of V1Reference
  * @throws IOException if the JSON string is invalid with respect to V1Reference
  */
  public static V1Reference fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, V1Reference.class);
  }

 /**
  * Convert an instance of V1Reference to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

