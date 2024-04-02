/*
 * Polyaxon SDKs and REST API specification.
 *    
 *
 * The version of the OpenAPI document: 2.1.5
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
import java.util.ArrayList;
import java.util.List;

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
 * V1Installation
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class V1Installation {
  public static final String SERIALIZED_NAME_KEY = "key";
  @SerializedName(SERIALIZED_NAME_KEY)
  private String key;

  public static final String SERIALIZED_NAME_VERSION = "version";
  @SerializedName(SERIALIZED_NAME_VERSION)
  private String version;

  public static final String SERIALIZED_NAME_DIST = "dist";
  @SerializedName(SERIALIZED_NAME_DIST)
  private String dist;

  public static final String SERIALIZED_NAME_HOST = "host";
  @SerializedName(SERIALIZED_NAME_HOST)
  private String host;

  public static final String SERIALIZED_NAME_HMAC = "hmac";
  @SerializedName(SERIALIZED_NAME_HMAC)
  private String hmac;

  public static final String SERIALIZED_NAME_MODE = "mode";
  @SerializedName(SERIALIZED_NAME_MODE)
  private String mode;

  public static final String SERIALIZED_NAME_ORGS = "orgs";
  @SerializedName(SERIALIZED_NAME_ORGS)
  private Boolean orgs;

  public static final String SERIALIZED_NAME_AUTH = "auth";
  @SerializedName(SERIALIZED_NAME_AUTH)
  private List<String> auth;

  public V1Installation() {
  }

  public V1Installation key(String key) {
    
    this.key = key;
    return this;
  }

   /**
   * Get key
   * @return key
  **/
  @javax.annotation.Nullable

  public String getKey() {
    return key;
  }


  public void setKey(String key) {
    this.key = key;
  }


  public V1Installation version(String version) {
    
    this.version = version;
    return this;
  }

   /**
   * Get version
   * @return version
  **/
  @javax.annotation.Nullable

  public String getVersion() {
    return version;
  }


  public void setVersion(String version) {
    this.version = version;
  }


  public V1Installation dist(String dist) {
    
    this.dist = dist;
    return this;
  }

   /**
   * Get dist
   * @return dist
  **/
  @javax.annotation.Nullable

  public String getDist() {
    return dist;
  }


  public void setDist(String dist) {
    this.dist = dist;
  }


  public V1Installation host(String host) {
    
    this.host = host;
    return this;
  }

   /**
   * Get host
   * @return host
  **/
  @javax.annotation.Nullable

  public String getHost() {
    return host;
  }


  public void setHost(String host) {
    this.host = host;
  }


  public V1Installation hmac(String hmac) {
    
    this.hmac = hmac;
    return this;
  }

   /**
   * Get hmac
   * @return hmac
  **/
  @javax.annotation.Nullable

  public String getHmac() {
    return hmac;
  }


  public void setHmac(String hmac) {
    this.hmac = hmac;
  }


  public V1Installation mode(String mode) {
    
    this.mode = mode;
    return this;
  }

   /**
   * Get mode
   * @return mode
  **/
  @javax.annotation.Nullable

  public String getMode() {
    return mode;
  }


  public void setMode(String mode) {
    this.mode = mode;
  }


  public V1Installation orgs(Boolean orgs) {
    
    this.orgs = orgs;
    return this;
  }

   /**
   * Get orgs
   * @return orgs
  **/
  @javax.annotation.Nullable

  public Boolean getOrgs() {
    return orgs;
  }


  public void setOrgs(Boolean orgs) {
    this.orgs = orgs;
  }


  public V1Installation auth(List<String> auth) {
    
    this.auth = auth;
    return this;
  }

  public V1Installation addAuthItem(String authItem) {
    if (this.auth == null) {
      this.auth = new ArrayList<>();
    }
    this.auth.add(authItem);
    return this;
  }

   /**
   * Get auth
   * @return auth
  **/
  @javax.annotation.Nullable

  public List<String> getAuth() {
    return auth;
  }


  public void setAuth(List<String> auth) {
    this.auth = auth;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1Installation v1Installation = (V1Installation) o;
    return Objects.equals(this.key, v1Installation.key) &&
        Objects.equals(this.version, v1Installation.version) &&
        Objects.equals(this.dist, v1Installation.dist) &&
        Objects.equals(this.host, v1Installation.host) &&
        Objects.equals(this.hmac, v1Installation.hmac) &&
        Objects.equals(this.mode, v1Installation.mode) &&
        Objects.equals(this.orgs, v1Installation.orgs) &&
        Objects.equals(this.auth, v1Installation.auth);
  }

  @Override
  public int hashCode() {
    return Objects.hash(key, version, dist, host, hmac, mode, orgs, auth);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1Installation {\n");
    sb.append("    key: ").append(toIndentedString(key)).append("\n");
    sb.append("    version: ").append(toIndentedString(version)).append("\n");
    sb.append("    dist: ").append(toIndentedString(dist)).append("\n");
    sb.append("    host: ").append(toIndentedString(host)).append("\n");
    sb.append("    hmac: ").append(toIndentedString(hmac)).append("\n");
    sb.append("    mode: ").append(toIndentedString(mode)).append("\n");
    sb.append("    orgs: ").append(toIndentedString(orgs)).append("\n");
    sb.append("    auth: ").append(toIndentedString(auth)).append("\n");
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
    openapiFields.add("key");
    openapiFields.add("version");
    openapiFields.add("dist");
    openapiFields.add("host");
    openapiFields.add("hmac");
    openapiFields.add("mode");
    openapiFields.add("orgs");
    openapiFields.add("auth");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to V1Installation
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!V1Installation.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in V1Installation is not found in the empty JSON string", V1Installation.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!V1Installation.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `V1Installation` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }
      if ((jsonObj.get("key") != null && !jsonObj.get("key").isJsonNull()) && !jsonObj.get("key").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `key` to be a primitive type in the JSON string but got `%s`", jsonObj.get("key").toString()));
      }
      if ((jsonObj.get("version") != null && !jsonObj.get("version").isJsonNull()) && !jsonObj.get("version").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `version` to be a primitive type in the JSON string but got `%s`", jsonObj.get("version").toString()));
      }
      if ((jsonObj.get("dist") != null && !jsonObj.get("dist").isJsonNull()) && !jsonObj.get("dist").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `dist` to be a primitive type in the JSON string but got `%s`", jsonObj.get("dist").toString()));
      }
      if ((jsonObj.get("host") != null && !jsonObj.get("host").isJsonNull()) && !jsonObj.get("host").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `host` to be a primitive type in the JSON string but got `%s`", jsonObj.get("host").toString()));
      }
      if ((jsonObj.get("hmac") != null && !jsonObj.get("hmac").isJsonNull()) && !jsonObj.get("hmac").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `hmac` to be a primitive type in the JSON string but got `%s`", jsonObj.get("hmac").toString()));
      }
      if ((jsonObj.get("mode") != null && !jsonObj.get("mode").isJsonNull()) && !jsonObj.get("mode").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `mode` to be a primitive type in the JSON string but got `%s`", jsonObj.get("mode").toString()));
      }
      // ensure the optional json data is an array if present
      if (jsonObj.get("auth") != null && !jsonObj.get("auth").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `auth` to be an array in the JSON string but got `%s`", jsonObj.get("auth").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!V1Installation.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'V1Installation' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<V1Installation> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(V1Installation.class));

       return (TypeAdapter<T>) new TypeAdapter<V1Installation>() {
           @Override
           public void write(JsonWriter out, V1Installation value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public V1Installation read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of V1Installation given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of V1Installation
  * @throws IOException if the JSON string is invalid with respect to V1Installation
  */
  public static V1Installation fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, V1Installation.class);
  }

 /**
  * Convert an instance of V1Installation to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

