/*
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0-rc22
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
import org.openapitools.client.model.SparkDeployMode;
import org.openapitools.client.model.V1SparkReplica;
import org.openapitools.client.model.V1SparkType;

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
 * V1Spark
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class V1Spark {
  public static final String SERIALIZED_NAME_KIND = "kind";
  @SerializedName(SERIALIZED_NAME_KIND)
  private String kind = "spark";

  public static final String SERIALIZED_NAME_CONNECTIONS = "connections";
  @SerializedName(SERIALIZED_NAME_CONNECTIONS)
  private List<String> connections;

  public static final String SERIALIZED_NAME_VOLUMES = "volumes";
  @SerializedName(SERIALIZED_NAME_VOLUMES)
  private List<Object> volumes;

  public static final String SERIALIZED_NAME_TYPE = "type";
  @SerializedName(SERIALIZED_NAME_TYPE)
  private V1SparkType type = V1SparkType.JAVA;

  public static final String SERIALIZED_NAME_SPARK_VERSION = "sparkVersion";
  @SerializedName(SERIALIZED_NAME_SPARK_VERSION)
  private String sparkVersion;

  public static final String SERIALIZED_NAME_PYTHON_VERSION = "pythonVersion";
  @SerializedName(SERIALIZED_NAME_PYTHON_VERSION)
  private String pythonVersion;

  public static final String SERIALIZED_NAME_DEPLOY_MODE = "deployMode";
  @SerializedName(SERIALIZED_NAME_DEPLOY_MODE)
  private SparkDeployMode deployMode = SparkDeployMode.CLUSTER;

  public static final String SERIALIZED_NAME_MAIN_CLASS = "mainClass";
  @SerializedName(SERIALIZED_NAME_MAIN_CLASS)
  private String mainClass;

  public static final String SERIALIZED_NAME_MAIN_APPLICATION_FILE = "mainApplicationFile";
  @SerializedName(SERIALIZED_NAME_MAIN_APPLICATION_FILE)
  private String mainApplicationFile;

  public static final String SERIALIZED_NAME_ARGUMENTS = "arguments";
  @SerializedName(SERIALIZED_NAME_ARGUMENTS)
  private List<String> arguments;

  public static final String SERIALIZED_NAME_HADOOP_CONF = "hadoopConf";
  @SerializedName(SERIALIZED_NAME_HADOOP_CONF)
  private Map<String, String> hadoopConf = new HashMap<>();

  public static final String SERIALIZED_NAME_SPARK_CONF = "sparkConf";
  @SerializedName(SERIALIZED_NAME_SPARK_CONF)
  private Map<String, String> sparkConf = new HashMap<>();

  public static final String SERIALIZED_NAME_SPARK_CONFIG_MAP = "sparkConfigMap";
  @SerializedName(SERIALIZED_NAME_SPARK_CONFIG_MAP)
  private String sparkConfigMap;

  public static final String SERIALIZED_NAME_HADOOP_CONFIG_MAP = "hadoopConfigMap";
  @SerializedName(SERIALIZED_NAME_HADOOP_CONFIG_MAP)
  private String hadoopConfigMap;

  public static final String SERIALIZED_NAME_EXECUTOR = "executor";
  @SerializedName(SERIALIZED_NAME_EXECUTOR)
  private V1SparkReplica executor;

  public static final String SERIALIZED_NAME_DRIVER = "driver";
  @SerializedName(SERIALIZED_NAME_DRIVER)
  private V1SparkReplica driver;

  public V1Spark() {
  }

  public V1Spark kind(String kind) {

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


  public V1Spark connections(List<String> connections) {

    this.connections = connections;
    return this;
  }

  public V1Spark addConnectionsItem(String connectionsItem) {
    if (this.connections == null) {
      this.connections = new ArrayList<>();
    }
    this.connections.add(connectionsItem);
    return this;
  }

   /**
   * Get connections
   * @return connections
  **/
  @javax.annotation.Nullable

  public List<String> getConnections() {
    return connections;
  }


  public void setConnections(List<String> connections) {
    this.connections = connections;
  }


  public V1Spark volumes(List<Object> volumes) {

    this.volumes = volumes;
    return this;
  }

  public V1Spark addVolumesItem(Object volumesItem) {
    if (this.volumes == null) {
      this.volumes = new ArrayList<>();
    }
    this.volumes.add(volumesItem);
    return this;
  }

   /**
   * Volumes is a list of volumes that can be mounted.
   * @return volumes
  **/
  @javax.annotation.Nullable

  public List<Object> getVolumes() {
    return volumes;
  }


  public void setVolumes(List<Object> volumes) {
    this.volumes = volumes;
  }


  public V1Spark type(V1SparkType type) {

    this.type = type;
    return this;
  }

   /**
   * Get type
   * @return type
  **/
  @javax.annotation.Nullable

  public V1SparkType getType() {
    return type;
  }


  public void setType(V1SparkType type) {
    this.type = type;
  }


  public V1Spark sparkVersion(String sparkVersion) {

    this.sparkVersion = sparkVersion;
    return this;
  }

   /**
   * Spark version is the version of Spark the application uses.
   * @return sparkVersion
  **/
  @javax.annotation.Nullable

  public String getSparkVersion() {
    return sparkVersion;
  }


  public void setSparkVersion(String sparkVersion) {
    this.sparkVersion = sparkVersion;
  }


  public V1Spark pythonVersion(String pythonVersion) {

    this.pythonVersion = pythonVersion;
    return this;
  }

   /**
   * Spark version is the version of Spark the application uses.
   * @return pythonVersion
  **/
  @javax.annotation.Nullable

  public String getPythonVersion() {
    return pythonVersion;
  }


  public void setPythonVersion(String pythonVersion) {
    this.pythonVersion = pythonVersion;
  }


  public V1Spark deployMode(SparkDeployMode deployMode) {

    this.deployMode = deployMode;
    return this;
  }

   /**
   * Get deployMode
   * @return deployMode
  **/
  @javax.annotation.Nullable

  public SparkDeployMode getDeployMode() {
    return deployMode;
  }


  public void setDeployMode(SparkDeployMode deployMode) {
    this.deployMode = deployMode;
  }


  public V1Spark mainClass(String mainClass) {

    this.mainClass = mainClass;
    return this;
  }

   /**
   * MainClass is the fully-qualified main class of the Spark application. This only applies to Java/Scala Spark applications.
   * @return mainClass
  **/
  @javax.annotation.Nullable

  public String getMainClass() {
    return mainClass;
  }


  public void setMainClass(String mainClass) {
    this.mainClass = mainClass;
  }


  public V1Spark mainApplicationFile(String mainApplicationFile) {

    this.mainApplicationFile = mainApplicationFile;
    return this;
  }

   /**
   * MainFile is the path to a bundled JAR, Python, or R file of the application.
   * @return mainApplicationFile
  **/
  @javax.annotation.Nullable

  public String getMainApplicationFile() {
    return mainApplicationFile;
  }


  public void setMainApplicationFile(String mainApplicationFile) {
    this.mainApplicationFile = mainApplicationFile;
  }


  public V1Spark arguments(List<String> arguments) {

    this.arguments = arguments;
    return this;
  }

  public V1Spark addArgumentsItem(String argumentsItem) {
    if (this.arguments == null) {
      this.arguments = new ArrayList<>();
    }
    this.arguments.add(argumentsItem);
    return this;
  }

   /**
   * Arguments is a list of arguments to be passed to the application.
   * @return arguments
  **/
  @javax.annotation.Nullable

  public List<String> getArguments() {
    return arguments;
  }


  public void setArguments(List<String> arguments) {
    this.arguments = arguments;
  }


  public V1Spark hadoopConf(Map<String, String> hadoopConf) {

    this.hadoopConf = hadoopConf;
    return this;
  }

  public V1Spark putHadoopConfItem(String key, String hadoopConfItem) {
    if (this.hadoopConf == null) {
      this.hadoopConf = new HashMap<>();
    }
    this.hadoopConf.put(key, hadoopConfItem);
    return this;
  }

   /**
   * HadoopConf carries user-specified Hadoop configuration properties as they would use the  the \&quot;--conf\&quot; option in spark-submit.  The SparkApplication controller automatically adds prefix \&quot;spark.hadoop.\&quot; to Hadoop configuration properties.
   * @return hadoopConf
  **/
  @javax.annotation.Nullable

  public Map<String, String> getHadoopConf() {
    return hadoopConf;
  }


  public void setHadoopConf(Map<String, String> hadoopConf) {
    this.hadoopConf = hadoopConf;
  }


  public V1Spark sparkConf(Map<String, String> sparkConf) {

    this.sparkConf = sparkConf;
    return this;
  }

  public V1Spark putSparkConfItem(String key, String sparkConfItem) {
    if (this.sparkConf == null) {
      this.sparkConf = new HashMap<>();
    }
    this.sparkConf.put(key, sparkConfItem);
    return this;
  }

   /**
   * HadoopConf carries user-specified Hadoop configuration properties as they would use the  the \&quot;--conf\&quot; option in spark-submit.  The SparkApplication controller automatically adds prefix \&quot;spark.hadoop.\&quot; to Hadoop configuration properties.
   * @return sparkConf
  **/
  @javax.annotation.Nullable

  public Map<String, String> getSparkConf() {
    return sparkConf;
  }


  public void setSparkConf(Map<String, String> sparkConf) {
    this.sparkConf = sparkConf;
  }


  public V1Spark sparkConfigMap(String sparkConfigMap) {

    this.sparkConfigMap = sparkConfigMap;
    return this;
  }

   /**
   * SparkConfigMap carries the name of the ConfigMap containing Spark configuration files such as log4j.properties. The controller will add environment variable SPARK_CONF_DIR to the path where the ConfigMap is mounted to.
   * @return sparkConfigMap
  **/
  @javax.annotation.Nullable

  public String getSparkConfigMap() {
    return sparkConfigMap;
  }


  public void setSparkConfigMap(String sparkConfigMap) {
    this.sparkConfigMap = sparkConfigMap;
  }


  public V1Spark hadoopConfigMap(String hadoopConfigMap) {

    this.hadoopConfigMap = hadoopConfigMap;
    return this;
  }

   /**
   * HadoopConfigMap carries the name of the ConfigMap containing Hadoop configuration files such as core-site.xml. The controller will add environment variable HADOOP_CONF_DIR to the path where the ConfigMap is mounted to.
   * @return hadoopConfigMap
  **/
  @javax.annotation.Nullable

  public String getHadoopConfigMap() {
    return hadoopConfigMap;
  }


  public void setHadoopConfigMap(String hadoopConfigMap) {
    this.hadoopConfigMap = hadoopConfigMap;
  }


  public V1Spark executor(V1SparkReplica executor) {

    this.executor = executor;
    return this;
  }

   /**
   * Get executor
   * @return executor
  **/
  @javax.annotation.Nullable

  public V1SparkReplica getExecutor() {
    return executor;
  }


  public void setExecutor(V1SparkReplica executor) {
    this.executor = executor;
  }


  public V1Spark driver(V1SparkReplica driver) {

    this.driver = driver;
    return this;
  }

   /**
   * Get driver
   * @return driver
  **/
  @javax.annotation.Nullable

  public V1SparkReplica getDriver() {
    return driver;
  }


  public void setDriver(V1SparkReplica driver) {
    this.driver = driver;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1Spark v1Spark = (V1Spark) o;
    return Objects.equals(this.kind, v1Spark.kind) &&
        Objects.equals(this.connections, v1Spark.connections) &&
        Objects.equals(this.volumes, v1Spark.volumes) &&
        Objects.equals(this.type, v1Spark.type) &&
        Objects.equals(this.sparkVersion, v1Spark.sparkVersion) &&
        Objects.equals(this.pythonVersion, v1Spark.pythonVersion) &&
        Objects.equals(this.deployMode, v1Spark.deployMode) &&
        Objects.equals(this.mainClass, v1Spark.mainClass) &&
        Objects.equals(this.mainApplicationFile, v1Spark.mainApplicationFile) &&
        Objects.equals(this.arguments, v1Spark.arguments) &&
        Objects.equals(this.hadoopConf, v1Spark.hadoopConf) &&
        Objects.equals(this.sparkConf, v1Spark.sparkConf) &&
        Objects.equals(this.sparkConfigMap, v1Spark.sparkConfigMap) &&
        Objects.equals(this.hadoopConfigMap, v1Spark.hadoopConfigMap) &&
        Objects.equals(this.executor, v1Spark.executor) &&
        Objects.equals(this.driver, v1Spark.driver);
  }

  @Override
  public int hashCode() {
    return Objects.hash(kind, connections, volumes, type, sparkVersion, pythonVersion, deployMode, mainClass, mainApplicationFile, arguments, hadoopConf, sparkConf, sparkConfigMap, hadoopConfigMap, executor, driver);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1Spark {\n");
    sb.append("    kind: ").append(toIndentedString(kind)).append("\n");
    sb.append("    connections: ").append(toIndentedString(connections)).append("\n");
    sb.append("    volumes: ").append(toIndentedString(volumes)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    sparkVersion: ").append(toIndentedString(sparkVersion)).append("\n");
    sb.append("    pythonVersion: ").append(toIndentedString(pythonVersion)).append("\n");
    sb.append("    deployMode: ").append(toIndentedString(deployMode)).append("\n");
    sb.append("    mainClass: ").append(toIndentedString(mainClass)).append("\n");
    sb.append("    mainApplicationFile: ").append(toIndentedString(mainApplicationFile)).append("\n");
    sb.append("    arguments: ").append(toIndentedString(arguments)).append("\n");
    sb.append("    hadoopConf: ").append(toIndentedString(hadoopConf)).append("\n");
    sb.append("    sparkConf: ").append(toIndentedString(sparkConf)).append("\n");
    sb.append("    sparkConfigMap: ").append(toIndentedString(sparkConfigMap)).append("\n");
    sb.append("    hadoopConfigMap: ").append(toIndentedString(hadoopConfigMap)).append("\n");
    sb.append("    executor: ").append(toIndentedString(executor)).append("\n");
    sb.append("    driver: ").append(toIndentedString(driver)).append("\n");
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
    openapiFields.add("connections");
    openapiFields.add("volumes");
    openapiFields.add("type");
    openapiFields.add("sparkVersion");
    openapiFields.add("pythonVersion");
    openapiFields.add("deployMode");
    openapiFields.add("mainClass");
    openapiFields.add("mainApplicationFile");
    openapiFields.add("arguments");
    openapiFields.add("hadoopConf");
    openapiFields.add("sparkConf");
    openapiFields.add("sparkConfigMap");
    openapiFields.add("hadoopConfigMap");
    openapiFields.add("executor");
    openapiFields.add("driver");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to V1Spark
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!V1Spark.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in V1Spark is not found in the empty JSON string", V1Spark.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!V1Spark.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `V1Spark` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }
      if ((jsonObj.get("kind") != null && !jsonObj.get("kind").isJsonNull()) && !jsonObj.get("kind").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `kind` to be a primitive type in the JSON string but got `%s`", jsonObj.get("kind").toString()));
      }
      // ensure the optional json data is an array if present
      if (jsonObj.get("connections") != null && !jsonObj.get("connections").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `connections` to be an array in the JSON string but got `%s`", jsonObj.get("connections").toString()));
      }
      // ensure the optional json data is an array if present
      if (jsonObj.get("volumes") != null && !jsonObj.get("volumes").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `volumes` to be an array in the JSON string but got `%s`", jsonObj.get("volumes").toString()));
      }
      if ((jsonObj.get("sparkVersion") != null && !jsonObj.get("sparkVersion").isJsonNull()) && !jsonObj.get("sparkVersion").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `sparkVersion` to be a primitive type in the JSON string but got `%s`", jsonObj.get("sparkVersion").toString()));
      }
      if ((jsonObj.get("pythonVersion") != null && !jsonObj.get("pythonVersion").isJsonNull()) && !jsonObj.get("pythonVersion").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `pythonVersion` to be a primitive type in the JSON string but got `%s`", jsonObj.get("pythonVersion").toString()));
      }
      if ((jsonObj.get("mainClass") != null && !jsonObj.get("mainClass").isJsonNull()) && !jsonObj.get("mainClass").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `mainClass` to be a primitive type in the JSON string but got `%s`", jsonObj.get("mainClass").toString()));
      }
      if ((jsonObj.get("mainApplicationFile") != null && !jsonObj.get("mainApplicationFile").isJsonNull()) && !jsonObj.get("mainApplicationFile").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `mainApplicationFile` to be a primitive type in the JSON string but got `%s`", jsonObj.get("mainApplicationFile").toString()));
      }
      // ensure the optional json data is an array if present
      if (jsonObj.get("arguments") != null && !jsonObj.get("arguments").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `arguments` to be an array in the JSON string but got `%s`", jsonObj.get("arguments").toString()));
      }
      if ((jsonObj.get("sparkConfigMap") != null && !jsonObj.get("sparkConfigMap").isJsonNull()) && !jsonObj.get("sparkConfigMap").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `sparkConfigMap` to be a primitive type in the JSON string but got `%s`", jsonObj.get("sparkConfigMap").toString()));
      }
      if ((jsonObj.get("hadoopConfigMap") != null && !jsonObj.get("hadoopConfigMap").isJsonNull()) && !jsonObj.get("hadoopConfigMap").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `hadoopConfigMap` to be a primitive type in the JSON string but got `%s`", jsonObj.get("hadoopConfigMap").toString()));
      }
      // validate the optional field `executor`
      if (jsonObj.get("executor") != null && !jsonObj.get("executor").isJsonNull()) {
        V1SparkReplica.validateJsonObject(jsonObj.getAsJsonObject("executor"));
      }
      // validate the optional field `driver`
      if (jsonObj.get("driver") != null && !jsonObj.get("driver").isJsonNull()) {
        V1SparkReplica.validateJsonObject(jsonObj.getAsJsonObject("driver"));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!V1Spark.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'V1Spark' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<V1Spark> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(V1Spark.class));

       return (TypeAdapter<T>) new TypeAdapter<V1Spark>() {
           @Override
           public void write(JsonWriter out, V1Spark value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public V1Spark read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of V1Spark given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of V1Spark
  * @throws IOException if the JSON string is invalid with respect to V1Spark
  */
  public static V1Spark fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, V1Spark.class);
  }

 /**
  * Convert an instance of V1Spark to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

