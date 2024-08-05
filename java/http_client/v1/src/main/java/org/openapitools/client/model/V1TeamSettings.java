/*
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.3.2
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
import org.openapitools.client.model.V1SettingsCatalog;

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
 * V1TeamSettings
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class V1TeamSettings {
  public static final String SERIALIZED_NAME_PROJECTS = "projects";
  @SerializedName(SERIALIZED_NAME_PROJECTS)
  private List<V1SettingsCatalog> projects;

  public static final String SERIALIZED_NAME_HUBS = "hubs";
  @SerializedName(SERIALIZED_NAME_HUBS)
  private List<V1SettingsCatalog> hubs;

  public static final String SERIALIZED_NAME_REGISTRIES = "registries";
  @SerializedName(SERIALIZED_NAME_REGISTRIES)
  private List<V1SettingsCatalog> registries;

  public V1TeamSettings() {
  }

  public V1TeamSettings projects(List<V1SettingsCatalog> projects) {

    this.projects = projects;
    return this;
  }

  public V1TeamSettings addProjectsItem(V1SettingsCatalog projectsItem) {
    if (this.projects == null) {
      this.projects = new ArrayList<>();
    }
    this.projects.add(projectsItem);
    return this;
  }

   /**
   * Get projects
   * @return projects
  **/
  @javax.annotation.Nullable

  public List<V1SettingsCatalog> getProjects() {
    return projects;
  }


  public void setProjects(List<V1SettingsCatalog> projects) {
    this.projects = projects;
  }


  public V1TeamSettings hubs(List<V1SettingsCatalog> hubs) {

    this.hubs = hubs;
    return this;
  }

  public V1TeamSettings addHubsItem(V1SettingsCatalog hubsItem) {
    if (this.hubs == null) {
      this.hubs = new ArrayList<>();
    }
    this.hubs.add(hubsItem);
    return this;
  }

   /**
   * Get hubs
   * @return hubs
  **/
  @javax.annotation.Nullable

  public List<V1SettingsCatalog> getHubs() {
    return hubs;
  }


  public void setHubs(List<V1SettingsCatalog> hubs) {
    this.hubs = hubs;
  }


  public V1TeamSettings registries(List<V1SettingsCatalog> registries) {

    this.registries = registries;
    return this;
  }

  public V1TeamSettings addRegistriesItem(V1SettingsCatalog registriesItem) {
    if (this.registries == null) {
      this.registries = new ArrayList<>();
    }
    this.registries.add(registriesItem);
    return this;
  }

   /**
   * Get registries
   * @return registries
  **/
  @javax.annotation.Nullable

  public List<V1SettingsCatalog> getRegistries() {
    return registries;
  }


  public void setRegistries(List<V1SettingsCatalog> registries) {
    this.registries = registries;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1TeamSettings v1TeamSettings = (V1TeamSettings) o;
    return Objects.equals(this.projects, v1TeamSettings.projects) &&
        Objects.equals(this.hubs, v1TeamSettings.hubs) &&
        Objects.equals(this.registries, v1TeamSettings.registries);
  }

  @Override
  public int hashCode() {
    return Objects.hash(projects, hubs, registries);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1TeamSettings {\n");
    sb.append("    projects: ").append(toIndentedString(projects)).append("\n");
    sb.append("    hubs: ").append(toIndentedString(hubs)).append("\n");
    sb.append("    registries: ").append(toIndentedString(registries)).append("\n");
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
    openapiFields.add("projects");
    openapiFields.add("hubs");
    openapiFields.add("registries");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to V1TeamSettings
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!V1TeamSettings.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in V1TeamSettings is not found in the empty JSON string", V1TeamSettings.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!V1TeamSettings.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `V1TeamSettings` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }
      if (jsonObj.get("projects") != null && !jsonObj.get("projects").isJsonNull()) {
        JsonArray jsonArrayprojects = jsonObj.getAsJsonArray("projects");
        if (jsonArrayprojects != null) {
          // ensure the json data is an array
          if (!jsonObj.get("projects").isJsonArray()) {
            throw new IllegalArgumentException(String.format("Expected the field `projects` to be an array in the JSON string but got `%s`", jsonObj.get("projects").toString()));
          }

          // validate the optional field `projects` (array)
          for (int i = 0; i < jsonArrayprojects.size(); i++) {
            V1SettingsCatalog.validateJsonObject(jsonArrayprojects.get(i).getAsJsonObject());
          };
        }
      }
      if (jsonObj.get("hubs") != null && !jsonObj.get("hubs").isJsonNull()) {
        JsonArray jsonArrayhubs = jsonObj.getAsJsonArray("hubs");
        if (jsonArrayhubs != null) {
          // ensure the json data is an array
          if (!jsonObj.get("hubs").isJsonArray()) {
            throw new IllegalArgumentException(String.format("Expected the field `hubs` to be an array in the JSON string but got `%s`", jsonObj.get("hubs").toString()));
          }

          // validate the optional field `hubs` (array)
          for (int i = 0; i < jsonArrayhubs.size(); i++) {
            V1SettingsCatalog.validateJsonObject(jsonArrayhubs.get(i).getAsJsonObject());
          };
        }
      }
      if (jsonObj.get("registries") != null && !jsonObj.get("registries").isJsonNull()) {
        JsonArray jsonArrayregistries = jsonObj.getAsJsonArray("registries");
        if (jsonArrayregistries != null) {
          // ensure the json data is an array
          if (!jsonObj.get("registries").isJsonArray()) {
            throw new IllegalArgumentException(String.format("Expected the field `registries` to be an array in the JSON string but got `%s`", jsonObj.get("registries").toString()));
          }

          // validate the optional field `registries` (array)
          for (int i = 0; i < jsonArrayregistries.size(); i++) {
            V1SettingsCatalog.validateJsonObject(jsonArrayregistries.get(i).getAsJsonObject());
          };
        }
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!V1TeamSettings.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'V1TeamSettings' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<V1TeamSettings> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(V1TeamSettings.class));

       return (TypeAdapter<T>) new TypeAdapter<V1TeamSettings>() {
           @Override
           public void write(JsonWriter out, V1TeamSettings value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public V1TeamSettings read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of V1TeamSettings given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of V1TeamSettings
  * @throws IOException if the JSON string is invalid with respect to V1TeamSettings
  */
  public static V1TeamSettings fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, V1TeamSettings.class);
  }

 /**
  * Convert an instance of V1TeamSettings to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

