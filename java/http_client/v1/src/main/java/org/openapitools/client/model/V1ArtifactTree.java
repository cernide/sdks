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
import java.util.HashMap;
import java.util.List;
import java.util.Map;

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
 * V1ArtifactTree
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class V1ArtifactTree {
  public static final String SERIALIZED_NAME_FILES = "files";
  @SerializedName(SERIALIZED_NAME_FILES)
  private Map<String, Integer> files = new HashMap<>();

  public static final String SERIALIZED_NAME_DIRS = "dirs";
  @SerializedName(SERIALIZED_NAME_DIRS)
  private List<String> dirs;

  public static final String SERIALIZED_NAME_IS_DONE = "is_done";
  @SerializedName(SERIALIZED_NAME_IS_DONE)
  private Boolean isDone;

  public V1ArtifactTree() {
  }

  public V1ArtifactTree files(Map<String, Integer> files) {
    
    this.files = files;
    return this;
  }

  public V1ArtifactTree putFilesItem(String key, Integer filesItem) {
    if (this.files == null) {
      this.files = new HashMap<>();
    }
    this.files.put(key, filesItem);
    return this;
  }

   /**
   * Get files
   * @return files
  **/
  @javax.annotation.Nullable

  public Map<String, Integer> getFiles() {
    return files;
  }


  public void setFiles(Map<String, Integer> files) {
    this.files = files;
  }


  public V1ArtifactTree dirs(List<String> dirs) {
    
    this.dirs = dirs;
    return this;
  }

  public V1ArtifactTree addDirsItem(String dirsItem) {
    if (this.dirs == null) {
      this.dirs = new ArrayList<>();
    }
    this.dirs.add(dirsItem);
    return this;
  }

   /**
   * Get dirs
   * @return dirs
  **/
  @javax.annotation.Nullable

  public List<String> getDirs() {
    return dirs;
  }


  public void setDirs(List<String> dirs) {
    this.dirs = dirs;
  }


  public V1ArtifactTree isDone(Boolean isDone) {
    
    this.isDone = isDone;
    return this;
  }

   /**
   * Get isDone
   * @return isDone
  **/
  @javax.annotation.Nullable

  public Boolean getIsDone() {
    return isDone;
  }


  public void setIsDone(Boolean isDone) {
    this.isDone = isDone;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1ArtifactTree v1ArtifactTree = (V1ArtifactTree) o;
    return Objects.equals(this.files, v1ArtifactTree.files) &&
        Objects.equals(this.dirs, v1ArtifactTree.dirs) &&
        Objects.equals(this.isDone, v1ArtifactTree.isDone);
  }

  @Override
  public int hashCode() {
    return Objects.hash(files, dirs, isDone);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1ArtifactTree {\n");
    sb.append("    files: ").append(toIndentedString(files)).append("\n");
    sb.append("    dirs: ").append(toIndentedString(dirs)).append("\n");
    sb.append("    isDone: ").append(toIndentedString(isDone)).append("\n");
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
    openapiFields.add("files");
    openapiFields.add("dirs");
    openapiFields.add("is_done");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to V1ArtifactTree
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!V1ArtifactTree.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in V1ArtifactTree is not found in the empty JSON string", V1ArtifactTree.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!V1ArtifactTree.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `V1ArtifactTree` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }
      // ensure the optional json data is an array if present
      if (jsonObj.get("dirs") != null && !jsonObj.get("dirs").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `dirs` to be an array in the JSON string but got `%s`", jsonObj.get("dirs").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!V1ArtifactTree.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'V1ArtifactTree' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<V1ArtifactTree> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(V1ArtifactTree.class));

       return (TypeAdapter<T>) new TypeAdapter<V1ArtifactTree>() {
           @Override
           public void write(JsonWriter out, V1ArtifactTree value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public V1ArtifactTree read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of V1ArtifactTree given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of V1ArtifactTree
  * @throws IOException if the JSON string is invalid with respect to V1ArtifactTree
  */
  public static V1ArtifactTree fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, V1ArtifactTree.class);
  }

 /**
  * Convert an instance of V1ArtifactTree to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

