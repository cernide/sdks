/*
 * Polyaxon SDKs and REST API specification.
 *    
 *
 * The version of the OpenAPI document: 2.0.0-rc14
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
 * V1DockerfileType
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class V1DockerfileType {
  public static final String SERIALIZED_NAME_IMAGE = "image";
  @SerializedName(SERIALIZED_NAME_IMAGE)
  private String image;

  public static final String SERIALIZED_NAME_ENV = "env";
  @SerializedName(SERIALIZED_NAME_ENV)
  private Map<String, String> env = new HashMap<>();

  public static final String SERIALIZED_NAME_PATH = "path";
  @SerializedName(SERIALIZED_NAME_PATH)
  private List<String> path;

  public static final String SERIALIZED_NAME_COPY = "copy";
  @SerializedName(SERIALIZED_NAME_COPY)
  private List<Object> copy;

  public static final String SERIALIZED_NAME_POST_RUN_COPY = "post_run_copy";
  @SerializedName(SERIALIZED_NAME_POST_RUN_COPY)
  private List<Object> postRunCopy;

  public static final String SERIALIZED_NAME_RUN = "run";
  @SerializedName(SERIALIZED_NAME_RUN)
  private List<String> run;

  public static final String SERIALIZED_NAME_LANG_ENV = "langEnv";
  @SerializedName(SERIALIZED_NAME_LANG_ENV)
  private String langEnv;

  public static final String SERIALIZED_NAME_UID = "uid";
  @SerializedName(SERIALIZED_NAME_UID)
  private Integer uid;

  public static final String SERIALIZED_NAME_GID = "gid";
  @SerializedName(SERIALIZED_NAME_GID)
  private Integer gid;

  public static final String SERIALIZED_NAME_USERNAME = "username";
  @SerializedName(SERIALIZED_NAME_USERNAME)
  private Integer username;

  public static final String SERIALIZED_NAME_FILENAME = "filename";
  @SerializedName(SERIALIZED_NAME_FILENAME)
  private String filename;

  public static final String SERIALIZED_NAME_WORKDIR = "workdir";
  @SerializedName(SERIALIZED_NAME_WORKDIR)
  private String workdir;

  public static final String SERIALIZED_NAME_WORKDIR_PATH = "workdirPath";
  @SerializedName(SERIALIZED_NAME_WORKDIR_PATH)
  private String workdirPath;

  public static final String SERIALIZED_NAME_SHELL = "shell";
  @SerializedName(SERIALIZED_NAME_SHELL)
  private String shell;

  public V1DockerfileType() {
  }

  public V1DockerfileType image(String image) {
    
    this.image = image;
    return this;
  }

   /**
   * Get image
   * @return image
  **/
  @javax.annotation.Nullable

  public String getImage() {
    return image;
  }


  public void setImage(String image) {
    this.image = image;
  }


  public V1DockerfileType env(Map<String, String> env) {
    
    this.env = env;
    return this;
  }

  public V1DockerfileType putEnvItem(String key, String envItem) {
    if (this.env == null) {
      this.env = new HashMap<>();
    }
    this.env.put(key, envItem);
    return this;
  }

   /**
   * Get env
   * @return env
  **/
  @javax.annotation.Nullable

  public Map<String, String> getEnv() {
    return env;
  }


  public void setEnv(Map<String, String> env) {
    this.env = env;
  }


  public V1DockerfileType path(List<String> path) {
    
    this.path = path;
    return this;
  }

  public V1DockerfileType addPathItem(String pathItem) {
    if (this.path == null) {
      this.path = new ArrayList<>();
    }
    this.path.add(pathItem);
    return this;
  }

   /**
   * Get path
   * @return path
  **/
  @javax.annotation.Nullable

  public List<String> getPath() {
    return path;
  }


  public void setPath(List<String> path) {
    this.path = path;
  }


  public V1DockerfileType copy(List<Object> copy) {
    
    this.copy = copy;
    return this;
  }

  public V1DockerfileType addCopyItem(Object copyItem) {
    if (this.copy == null) {
      this.copy = new ArrayList<>();
    }
    this.copy.add(copyItem);
    return this;
  }

   /**
   * Get copy
   * @return copy
  **/
  @javax.annotation.Nullable

  public List<Object> getCopy() {
    return copy;
  }


  public void setCopy(List<Object> copy) {
    this.copy = copy;
  }


  public V1DockerfileType postRunCopy(List<Object> postRunCopy) {
    
    this.postRunCopy = postRunCopy;
    return this;
  }

  public V1DockerfileType addPostRunCopyItem(Object postRunCopyItem) {
    if (this.postRunCopy == null) {
      this.postRunCopy = new ArrayList<>();
    }
    this.postRunCopy.add(postRunCopyItem);
    return this;
  }

   /**
   * Get postRunCopy
   * @return postRunCopy
  **/
  @javax.annotation.Nullable

  public List<Object> getPostRunCopy() {
    return postRunCopy;
  }


  public void setPostRunCopy(List<Object> postRunCopy) {
    this.postRunCopy = postRunCopy;
  }


  public V1DockerfileType run(List<String> run) {
    
    this.run = run;
    return this;
  }

  public V1DockerfileType addRunItem(String runItem) {
    if (this.run == null) {
      this.run = new ArrayList<>();
    }
    this.run.add(runItem);
    return this;
  }

   /**
   * Get run
   * @return run
  **/
  @javax.annotation.Nullable

  public List<String> getRun() {
    return run;
  }


  public void setRun(List<String> run) {
    this.run = run;
  }


  public V1DockerfileType langEnv(String langEnv) {
    
    this.langEnv = langEnv;
    return this;
  }

   /**
   * Get langEnv
   * @return langEnv
  **/
  @javax.annotation.Nullable

  public String getLangEnv() {
    return langEnv;
  }


  public void setLangEnv(String langEnv) {
    this.langEnv = langEnv;
  }


  public V1DockerfileType uid(Integer uid) {
    
    this.uid = uid;
    return this;
  }

   /**
   * Get uid
   * @return uid
  **/
  @javax.annotation.Nullable

  public Integer getUid() {
    return uid;
  }


  public void setUid(Integer uid) {
    this.uid = uid;
  }


  public V1DockerfileType gid(Integer gid) {
    
    this.gid = gid;
    return this;
  }

   /**
   * Get gid
   * @return gid
  **/
  @javax.annotation.Nullable

  public Integer getGid() {
    return gid;
  }


  public void setGid(Integer gid) {
    this.gid = gid;
  }


  public V1DockerfileType username(Integer username) {
    
    this.username = username;
    return this;
  }

   /**
   * Get username
   * @return username
  **/
  @javax.annotation.Nullable

  public Integer getUsername() {
    return username;
  }


  public void setUsername(Integer username) {
    this.username = username;
  }


  public V1DockerfileType filename(String filename) {
    
    this.filename = filename;
    return this;
  }

   /**
   * Get filename
   * @return filename
  **/
  @javax.annotation.Nullable

  public String getFilename() {
    return filename;
  }


  public void setFilename(String filename) {
    this.filename = filename;
  }


  public V1DockerfileType workdir(String workdir) {
    
    this.workdir = workdir;
    return this;
  }

   /**
   * Get workdir
   * @return workdir
  **/
  @javax.annotation.Nullable

  public String getWorkdir() {
    return workdir;
  }


  public void setWorkdir(String workdir) {
    this.workdir = workdir;
  }


  public V1DockerfileType workdirPath(String workdirPath) {
    
    this.workdirPath = workdirPath;
    return this;
  }

   /**
   * Get workdirPath
   * @return workdirPath
  **/
  @javax.annotation.Nullable

  public String getWorkdirPath() {
    return workdirPath;
  }


  public void setWorkdirPath(String workdirPath) {
    this.workdirPath = workdirPath;
  }


  public V1DockerfileType shell(String shell) {
    
    this.shell = shell;
    return this;
  }

   /**
   * Get shell
   * @return shell
  **/
  @javax.annotation.Nullable

  public String getShell() {
    return shell;
  }


  public void setShell(String shell) {
    this.shell = shell;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1DockerfileType v1DockerfileType = (V1DockerfileType) o;
    return Objects.equals(this.image, v1DockerfileType.image) &&
        Objects.equals(this.env, v1DockerfileType.env) &&
        Objects.equals(this.path, v1DockerfileType.path) &&
        Objects.equals(this.copy, v1DockerfileType.copy) &&
        Objects.equals(this.postRunCopy, v1DockerfileType.postRunCopy) &&
        Objects.equals(this.run, v1DockerfileType.run) &&
        Objects.equals(this.langEnv, v1DockerfileType.langEnv) &&
        Objects.equals(this.uid, v1DockerfileType.uid) &&
        Objects.equals(this.gid, v1DockerfileType.gid) &&
        Objects.equals(this.username, v1DockerfileType.username) &&
        Objects.equals(this.filename, v1DockerfileType.filename) &&
        Objects.equals(this.workdir, v1DockerfileType.workdir) &&
        Objects.equals(this.workdirPath, v1DockerfileType.workdirPath) &&
        Objects.equals(this.shell, v1DockerfileType.shell);
  }

  @Override
  public int hashCode() {
    return Objects.hash(image, env, path, copy, postRunCopy, run, langEnv, uid, gid, username, filename, workdir, workdirPath, shell);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1DockerfileType {\n");
    sb.append("    image: ").append(toIndentedString(image)).append("\n");
    sb.append("    env: ").append(toIndentedString(env)).append("\n");
    sb.append("    path: ").append(toIndentedString(path)).append("\n");
    sb.append("    copy: ").append(toIndentedString(copy)).append("\n");
    sb.append("    postRunCopy: ").append(toIndentedString(postRunCopy)).append("\n");
    sb.append("    run: ").append(toIndentedString(run)).append("\n");
    sb.append("    langEnv: ").append(toIndentedString(langEnv)).append("\n");
    sb.append("    uid: ").append(toIndentedString(uid)).append("\n");
    sb.append("    gid: ").append(toIndentedString(gid)).append("\n");
    sb.append("    username: ").append(toIndentedString(username)).append("\n");
    sb.append("    filename: ").append(toIndentedString(filename)).append("\n");
    sb.append("    workdir: ").append(toIndentedString(workdir)).append("\n");
    sb.append("    workdirPath: ").append(toIndentedString(workdirPath)).append("\n");
    sb.append("    shell: ").append(toIndentedString(shell)).append("\n");
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
    openapiFields.add("image");
    openapiFields.add("env");
    openapiFields.add("path");
    openapiFields.add("copy");
    openapiFields.add("post_run_copy");
    openapiFields.add("run");
    openapiFields.add("langEnv");
    openapiFields.add("uid");
    openapiFields.add("gid");
    openapiFields.add("username");
    openapiFields.add("filename");
    openapiFields.add("workdir");
    openapiFields.add("workdirPath");
    openapiFields.add("shell");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to V1DockerfileType
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!V1DockerfileType.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in V1DockerfileType is not found in the empty JSON string", V1DockerfileType.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!V1DockerfileType.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `V1DockerfileType` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }
      if ((jsonObj.get("image") != null && !jsonObj.get("image").isJsonNull()) && !jsonObj.get("image").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `image` to be a primitive type in the JSON string but got `%s`", jsonObj.get("image").toString()));
      }
      // ensure the optional json data is an array if present
      if (jsonObj.get("path") != null && !jsonObj.get("path").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `path` to be an array in the JSON string but got `%s`", jsonObj.get("path").toString()));
      }
      // ensure the optional json data is an array if present
      if (jsonObj.get("copy") != null && !jsonObj.get("copy").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `copy` to be an array in the JSON string but got `%s`", jsonObj.get("copy").toString()));
      }
      // ensure the optional json data is an array if present
      if (jsonObj.get("post_run_copy") != null && !jsonObj.get("post_run_copy").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `post_run_copy` to be an array in the JSON string but got `%s`", jsonObj.get("post_run_copy").toString()));
      }
      // ensure the optional json data is an array if present
      if (jsonObj.get("run") != null && !jsonObj.get("run").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `run` to be an array in the JSON string but got `%s`", jsonObj.get("run").toString()));
      }
      if ((jsonObj.get("langEnv") != null && !jsonObj.get("langEnv").isJsonNull()) && !jsonObj.get("langEnv").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `langEnv` to be a primitive type in the JSON string but got `%s`", jsonObj.get("langEnv").toString()));
      }
      if ((jsonObj.get("filename") != null && !jsonObj.get("filename").isJsonNull()) && !jsonObj.get("filename").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `filename` to be a primitive type in the JSON string but got `%s`", jsonObj.get("filename").toString()));
      }
      if ((jsonObj.get("workdir") != null && !jsonObj.get("workdir").isJsonNull()) && !jsonObj.get("workdir").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `workdir` to be a primitive type in the JSON string but got `%s`", jsonObj.get("workdir").toString()));
      }
      if ((jsonObj.get("workdirPath") != null && !jsonObj.get("workdirPath").isJsonNull()) && !jsonObj.get("workdirPath").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `workdirPath` to be a primitive type in the JSON string but got `%s`", jsonObj.get("workdirPath").toString()));
      }
      if ((jsonObj.get("shell") != null && !jsonObj.get("shell").isJsonNull()) && !jsonObj.get("shell").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `shell` to be a primitive type in the JSON string but got `%s`", jsonObj.get("shell").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!V1DockerfileType.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'V1DockerfileType' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<V1DockerfileType> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(V1DockerfileType.class));

       return (TypeAdapter<T>) new TypeAdapter<V1DockerfileType>() {
           @Override
           public void write(JsonWriter out, V1DockerfileType value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public V1DockerfileType read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of V1DockerfileType given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of V1DockerfileType
  * @throws IOException if the JSON string is invalid with respect to V1DockerfileType
  */
  public static V1DockerfileType fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, V1DockerfileType.class);
  }

 /**
  * Convert an instance of V1DockerfileType to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

