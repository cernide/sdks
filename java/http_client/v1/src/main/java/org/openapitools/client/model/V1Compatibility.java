/*
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.4.2
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
import org.openapitools.client.model.V1Version;

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
 * V1Compatibility
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class V1Compatibility {
  public static final String SERIALIZED_NAME_CLI = "cli";
  @SerializedName(SERIALIZED_NAME_CLI)
  private V1Version cli;

  public static final String SERIALIZED_NAME_PLATFORM = "platform";
  @SerializedName(SERIALIZED_NAME_PLATFORM)
  private V1Version platform;

  public static final String SERIALIZED_NAME_AGENT = "agent";
  @SerializedName(SERIALIZED_NAME_AGENT)
  private V1Version agent;

  public static final String SERIALIZED_NAME_UI = "ui";
  @SerializedName(SERIALIZED_NAME_UI)
  private V1Version ui;

  public V1Compatibility() {
  }

  public V1Compatibility cli(V1Version cli) {

    this.cli = cli;
    return this;
  }

   /**
   * Get cli
   * @return cli
  **/
  @javax.annotation.Nullable

  public V1Version getCli() {
    return cli;
  }


  public void setCli(V1Version cli) {
    this.cli = cli;
  }


  public V1Compatibility platform(V1Version platform) {

    this.platform = platform;
    return this;
  }

   /**
   * Get platform
   * @return platform
  **/
  @javax.annotation.Nullable

  public V1Version getPlatform() {
    return platform;
  }


  public void setPlatform(V1Version platform) {
    this.platform = platform;
  }


  public V1Compatibility agent(V1Version agent) {

    this.agent = agent;
    return this;
  }

   /**
   * Get agent
   * @return agent
  **/
  @javax.annotation.Nullable

  public V1Version getAgent() {
    return agent;
  }


  public void setAgent(V1Version agent) {
    this.agent = agent;
  }


  public V1Compatibility ui(V1Version ui) {

    this.ui = ui;
    return this;
  }

   /**
   * Get ui
   * @return ui
  **/
  @javax.annotation.Nullable

  public V1Version getUi() {
    return ui;
  }


  public void setUi(V1Version ui) {
    this.ui = ui;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1Compatibility v1Compatibility = (V1Compatibility) o;
    return Objects.equals(this.cli, v1Compatibility.cli) &&
        Objects.equals(this.platform, v1Compatibility.platform) &&
        Objects.equals(this.agent, v1Compatibility.agent) &&
        Objects.equals(this.ui, v1Compatibility.ui);
  }

  @Override
  public int hashCode() {
    return Objects.hash(cli, platform, agent, ui);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1Compatibility {\n");
    sb.append("    cli: ").append(toIndentedString(cli)).append("\n");
    sb.append("    platform: ").append(toIndentedString(platform)).append("\n");
    sb.append("    agent: ").append(toIndentedString(agent)).append("\n");
    sb.append("    ui: ").append(toIndentedString(ui)).append("\n");
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
    openapiFields.add("cli");
    openapiFields.add("platform");
    openapiFields.add("agent");
    openapiFields.add("ui");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to V1Compatibility
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!V1Compatibility.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in V1Compatibility is not found in the empty JSON string", V1Compatibility.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!V1Compatibility.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `V1Compatibility` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }
      // validate the optional field `cli`
      if (jsonObj.get("cli") != null && !jsonObj.get("cli").isJsonNull()) {
        V1Version.validateJsonObject(jsonObj.getAsJsonObject("cli"));
      }
      // validate the optional field `platform`
      if (jsonObj.get("platform") != null && !jsonObj.get("platform").isJsonNull()) {
        V1Version.validateJsonObject(jsonObj.getAsJsonObject("platform"));
      }
      // validate the optional field `agent`
      if (jsonObj.get("agent") != null && !jsonObj.get("agent").isJsonNull()) {
        V1Version.validateJsonObject(jsonObj.getAsJsonObject("agent"));
      }
      // validate the optional field `ui`
      if (jsonObj.get("ui") != null && !jsonObj.get("ui").isJsonNull()) {
        V1Version.validateJsonObject(jsonObj.getAsJsonObject("ui"));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!V1Compatibility.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'V1Compatibility' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<V1Compatibility> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(V1Compatibility.class));

       return (TypeAdapter<T>) new TypeAdapter<V1Compatibility>() {
           @Override
           public void write(JsonWriter out, V1Compatibility value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public V1Compatibility read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of V1Compatibility given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of V1Compatibility
  * @throws IOException if the JSON string is invalid with respect to V1Compatibility
  */
  public static V1Compatibility fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, V1Compatibility.class);
  }

 /**
  * Convert an instance of V1Compatibility to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

