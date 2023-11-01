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
import java.time.OffsetDateTime;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.client.model.V1TeamSettings;

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
 * V1Team
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class V1Team {
  public static final String SERIALIZED_NAME_UUID = "uuid";
  @SerializedName(SERIALIZED_NAME_UUID)
  private String uuid;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_PROJECTS = "projects";
  @SerializedName(SERIALIZED_NAME_PROJECTS)
  private List<String> projects;

  public static final String SERIALIZED_NAME_COMPONENT_HUBS = "component_hubs";
  @SerializedName(SERIALIZED_NAME_COMPONENT_HUBS)
  private List<String> componentHubs;

  public static final String SERIALIZED_NAME_MODEL_REGISTRIES = "model_registries";
  @SerializedName(SERIALIZED_NAME_MODEL_REGISTRIES)
  private List<String> modelRegistries;

  public static final String SERIALIZED_NAME_SETTINGS = "settings";
  @SerializedName(SERIALIZED_NAME_SETTINGS)
  private V1TeamSettings settings;

  public static final String SERIALIZED_NAME_CREATED_AT = "created_at";
  @SerializedName(SERIALIZED_NAME_CREATED_AT)
  private OffsetDateTime createdAt;

  public static final String SERIALIZED_NAME_UPDATED_AT = "updated_at";
  @SerializedName(SERIALIZED_NAME_UPDATED_AT)
  private OffsetDateTime updatedAt;

  public V1Team() {
  }

  public V1Team uuid(String uuid) {

    this.uuid = uuid;
    return this;
  }

   /**
   * Get uuid
   * @return uuid
  **/
  @javax.annotation.Nullable

  public String getUuid() {
    return uuid;
  }


  public void setUuid(String uuid) {
    this.uuid = uuid;
  }


  public V1Team name(String name) {

    this.name = name;
    return this;
  }

   /**
   * Get name
   * @return name
  **/
  @javax.annotation.Nullable

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public V1Team projects(List<String> projects) {

    this.projects = projects;
    return this;
  }

  public V1Team addProjectsItem(String projectsItem) {
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

  public List<String> getProjects() {
    return projects;
  }


  public void setProjects(List<String> projects) {
    this.projects = projects;
  }


  public V1Team componentHubs(List<String> componentHubs) {

    this.componentHubs = componentHubs;
    return this;
  }

  public V1Team addComponentHubsItem(String componentHubsItem) {
    if (this.componentHubs == null) {
      this.componentHubs = new ArrayList<>();
    }
    this.componentHubs.add(componentHubsItem);
    return this;
  }

   /**
   * Get componentHubs
   * @return componentHubs
  **/
  @javax.annotation.Nullable

  public List<String> getComponentHubs() {
    return componentHubs;
  }


  public void setComponentHubs(List<String> componentHubs) {
    this.componentHubs = componentHubs;
  }


  public V1Team modelRegistries(List<String> modelRegistries) {

    this.modelRegistries = modelRegistries;
    return this;
  }

  public V1Team addModelRegistriesItem(String modelRegistriesItem) {
    if (this.modelRegistries == null) {
      this.modelRegistries = new ArrayList<>();
    }
    this.modelRegistries.add(modelRegistriesItem);
    return this;
  }

   /**
   * Get modelRegistries
   * @return modelRegistries
  **/
  @javax.annotation.Nullable

  public List<String> getModelRegistries() {
    return modelRegistries;
  }


  public void setModelRegistries(List<String> modelRegistries) {
    this.modelRegistries = modelRegistries;
  }


  public V1Team settings(V1TeamSettings settings) {

    this.settings = settings;
    return this;
  }

   /**
   * Get settings
   * @return settings
  **/
  @javax.annotation.Nullable

  public V1TeamSettings getSettings() {
    return settings;
  }


  public void setSettings(V1TeamSettings settings) {
    this.settings = settings;
  }


  public V1Team createdAt(OffsetDateTime createdAt) {

    this.createdAt = createdAt;
    return this;
  }

   /**
   * Get createdAt
   * @return createdAt
  **/
  @javax.annotation.Nullable

  public OffsetDateTime getCreatedAt() {
    return createdAt;
  }


  public void setCreatedAt(OffsetDateTime createdAt) {
    this.createdAt = createdAt;
  }


  public V1Team updatedAt(OffsetDateTime updatedAt) {

    this.updatedAt = updatedAt;
    return this;
  }

   /**
   * Get updatedAt
   * @return updatedAt
  **/
  @javax.annotation.Nullable

  public OffsetDateTime getUpdatedAt() {
    return updatedAt;
  }


  public void setUpdatedAt(OffsetDateTime updatedAt) {
    this.updatedAt = updatedAt;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1Team v1Team = (V1Team) o;
    return Objects.equals(this.uuid, v1Team.uuid) &&
        Objects.equals(this.name, v1Team.name) &&
        Objects.equals(this.projects, v1Team.projects) &&
        Objects.equals(this.componentHubs, v1Team.componentHubs) &&
        Objects.equals(this.modelRegistries, v1Team.modelRegistries) &&
        Objects.equals(this.settings, v1Team.settings) &&
        Objects.equals(this.createdAt, v1Team.createdAt) &&
        Objects.equals(this.updatedAt, v1Team.updatedAt);
  }

  @Override
  public int hashCode() {
    return Objects.hash(uuid, name, projects, componentHubs, modelRegistries, settings, createdAt, updatedAt);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1Team {\n");
    sb.append("    uuid: ").append(toIndentedString(uuid)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    projects: ").append(toIndentedString(projects)).append("\n");
    sb.append("    componentHubs: ").append(toIndentedString(componentHubs)).append("\n");
    sb.append("    modelRegistries: ").append(toIndentedString(modelRegistries)).append("\n");
    sb.append("    settings: ").append(toIndentedString(settings)).append("\n");
    sb.append("    createdAt: ").append(toIndentedString(createdAt)).append("\n");
    sb.append("    updatedAt: ").append(toIndentedString(updatedAt)).append("\n");
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
    openapiFields.add("uuid");
    openapiFields.add("name");
    openapiFields.add("projects");
    openapiFields.add("component_hubs");
    openapiFields.add("model_registries");
    openapiFields.add("settings");
    openapiFields.add("created_at");
    openapiFields.add("updated_at");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to V1Team
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!V1Team.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in V1Team is not found in the empty JSON string", V1Team.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!V1Team.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `V1Team` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }
      if ((jsonObj.get("uuid") != null && !jsonObj.get("uuid").isJsonNull()) && !jsonObj.get("uuid").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `uuid` to be a primitive type in the JSON string but got `%s`", jsonObj.get("uuid").toString()));
      }
      if ((jsonObj.get("name") != null && !jsonObj.get("name").isJsonNull()) && !jsonObj.get("name").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `name` to be a primitive type in the JSON string but got `%s`", jsonObj.get("name").toString()));
      }
      // ensure the optional json data is an array if present
      if (jsonObj.get("projects") != null && !jsonObj.get("projects").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `projects` to be an array in the JSON string but got `%s`", jsonObj.get("projects").toString()));
      }
      // ensure the optional json data is an array if present
      if (jsonObj.get("component_hubs") != null && !jsonObj.get("component_hubs").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `component_hubs` to be an array in the JSON string but got `%s`", jsonObj.get("component_hubs").toString()));
      }
      // ensure the optional json data is an array if present
      if (jsonObj.get("model_registries") != null && !jsonObj.get("model_registries").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `model_registries` to be an array in the JSON string but got `%s`", jsonObj.get("model_registries").toString()));
      }
      // validate the optional field `settings`
      if (jsonObj.get("settings") != null && !jsonObj.get("settings").isJsonNull()) {
        V1TeamSettings.validateJsonObject(jsonObj.getAsJsonObject("settings"));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!V1Team.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'V1Team' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<V1Team> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(V1Team.class));

       return (TypeAdapter<T>) new TypeAdapter<V1Team>() {
           @Override
           public void write(JsonWriter out, V1Team value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public V1Team read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of V1Team given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of V1Team
  * @throws IOException if the JSON string is invalid with respect to V1Team
  */
  public static V1Team fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, V1Team.class);
  }

 /**
  * Convert an instance of V1Team to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

