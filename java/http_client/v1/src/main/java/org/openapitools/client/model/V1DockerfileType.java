// Copyright 2018-2021 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * The version of the OpenAPI document: 1.17.1
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
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

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
  private Map<String, String> env = null;

  public static final String SERIALIZED_NAME_PATH = "path";
  @SerializedName(SERIALIZED_NAME_PATH)
  private List<String> path = null;

  public static final String SERIALIZED_NAME_COPY = "copy";
  @SerializedName(SERIALIZED_NAME_COPY)
  private List<Object> copy = null;

  public static final String SERIALIZED_NAME_POST_RUN_COPY = "post_run_copy";
  @SerializedName(SERIALIZED_NAME_POST_RUN_COPY)
  private List<Object> postRunCopy = null;

  public static final String SERIALIZED_NAME_RUN = "run";
  @SerializedName(SERIALIZED_NAME_RUN)
  private List<String> run = null;

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


  public V1DockerfileType image(String image) {
    
    this.image = image;
    return this;
  }

   /**
   * Get image
   * @return image
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

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
      this.env = new HashMap<String, String>();
    }
    this.env.put(key, envItem);
    return this;
  }

   /**
   * Get env
   * @return env
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

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
      this.path = new ArrayList<String>();
    }
    this.path.add(pathItem);
    return this;
  }

   /**
   * Get path
   * @return path
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

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
      this.copy = new ArrayList<Object>();
    }
    this.copy.add(copyItem);
    return this;
  }

   /**
   * Get copy
   * @return copy
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

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
      this.postRunCopy = new ArrayList<Object>();
    }
    this.postRunCopy.add(postRunCopyItem);
    return this;
  }

   /**
   * Get postRunCopy
   * @return postRunCopy
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

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
      this.run = new ArrayList<String>();
    }
    this.run.add(runItem);
    return this;
  }

   /**
   * Get run
   * @return run
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

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
  @ApiModelProperty(value = "")

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
  @ApiModelProperty(value = "")

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
  @ApiModelProperty(value = "")

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
  @ApiModelProperty(value = "")

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
  @ApiModelProperty(value = "")

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
  @ApiModelProperty(value = "")

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
  @ApiModelProperty(value = "")

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
  @ApiModelProperty(value = "")

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

}

