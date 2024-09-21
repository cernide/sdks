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
import org.openapitools.client.model.V1ArtifactsMount;
import org.openapitools.client.model.V1ArtifactsType;
import org.openapitools.client.model.V1AuthType;
import org.openapitools.client.model.V1CompiledOperation;
import org.openapitools.client.model.V1ConnectionResource;
import org.openapitools.client.model.V1ConnectionSchema;
import org.openapitools.client.model.V1ConnectionType;
import org.openapitools.client.model.V1EarlyStopping;
import org.openapitools.client.model.V1Event;
import org.openapitools.client.model.V1EventType;
import org.openapitools.client.model.V1GcsType;
import org.openapitools.client.model.V1HpParams;
import org.openapitools.client.model.V1Matrix;
import org.openapitools.client.model.V1MatrixKind;
import org.openapitools.client.model.V1Operation;
import org.openapitools.client.model.V1PolyaxonInitContainer;
import org.openapitools.client.model.V1PolyaxonSidecarContainer;
import org.openapitools.client.model.V1Reference;
import org.openapitools.client.model.V1RunSchema;
import org.openapitools.client.model.V1S3Type;
import org.openapitools.client.model.V1Schedule;
import org.openapitools.client.model.V1ScheduleKind;
import org.openapitools.client.model.V1UriType;
import org.openapitools.client.model.V1WasbType;

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
 * V1Schemas
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class V1Schemas {
  public static final String SERIALIZED_NAME_EARLY_STOPPING = "earlyStopping";
  @SerializedName(SERIALIZED_NAME_EARLY_STOPPING)
  private V1EarlyStopping earlyStopping;

  public static final String SERIALIZED_NAME_MATRIX = "matrix";
  @SerializedName(SERIALIZED_NAME_MATRIX)
  private V1Matrix matrix;

  public static final String SERIALIZED_NAME_RUN = "run";
  @SerializedName(SERIALIZED_NAME_RUN)
  private V1RunSchema run;

  public static final String SERIALIZED_NAME_OPERATION = "operation";
  @SerializedName(SERIALIZED_NAME_OPERATION)
  private V1Operation operation;

  public static final String SERIALIZED_NAME_COMPILED_OPERATION = "compiledOperation";
  @SerializedName(SERIALIZED_NAME_COMPILED_OPERATION)
  private V1CompiledOperation compiledOperation;

  public static final String SERIALIZED_NAME_SCHEDULE = "schedule";
  @SerializedName(SERIALIZED_NAME_SCHEDULE)
  private V1Schedule schedule;

  public static final String SERIALIZED_NAME_CONNECTION_SCHEMA = "connectionSchema";
  @SerializedName(SERIALIZED_NAME_CONNECTION_SCHEMA)
  private V1ConnectionSchema connectionSchema;

  public static final String SERIALIZED_NAME_HP_PARAMS = "hpParams";
  @SerializedName(SERIALIZED_NAME_HP_PARAMS)
  private V1HpParams hpParams;

  public static final String SERIALIZED_NAME_REFERENCE = "reference";
  @SerializedName(SERIALIZED_NAME_REFERENCE)
  private V1Reference reference;

  public static final String SERIALIZED_NAME_ARTIFACTS_MOUNT = "artifactsMount";
  @SerializedName(SERIALIZED_NAME_ARTIFACTS_MOUNT)
  private V1ArtifactsMount artifactsMount;

  public static final String SERIALIZED_NAME_POLYAXON_SIDECAR_CONTAINER = "polyaxonSidecarContainer";
  @SerializedName(SERIALIZED_NAME_POLYAXON_SIDECAR_CONTAINER)
  private V1PolyaxonSidecarContainer polyaxonSidecarContainer;

  public static final String SERIALIZED_NAME_POLYAXON_INIT_CONTAINER = "polyaxonInitContainer";
  @SerializedName(SERIALIZED_NAME_POLYAXON_INIT_CONTAINER)
  private V1PolyaxonInitContainer polyaxonInitContainer;

  public static final String SERIALIZED_NAME_ARTIFACS = "artifacs";
  @SerializedName(SERIALIZED_NAME_ARTIFACS)
  private V1ArtifactsType artifacs;

  public static final String SERIALIZED_NAME_WASB = "wasb";
  @SerializedName(SERIALIZED_NAME_WASB)
  private V1WasbType wasb;

  public static final String SERIALIZED_NAME_GCS = "gcs";
  @SerializedName(SERIALIZED_NAME_GCS)
  private V1GcsType gcs;

  public static final String SERIALIZED_NAME_S3 = "s3";
  @SerializedName(SERIALIZED_NAME_S3)
  private V1S3Type s3;

  public static final String SERIALIZED_NAME_AUTH = "auth";
  @SerializedName(SERIALIZED_NAME_AUTH)
  private V1AuthType auth;

  public static final String SERIALIZED_NAME_URI = "uri";
  @SerializedName(SERIALIZED_NAME_URI)
  private V1UriType uri;

  public static final String SERIALIZED_NAME_RESOURCE = "resource";
  @SerializedName(SERIALIZED_NAME_RESOURCE)
  private V1ConnectionResource resource;

  public static final String SERIALIZED_NAME_CONNECTION = "connection";
  @SerializedName(SERIALIZED_NAME_CONNECTION)
  private V1ConnectionType connection;

  public static final String SERIALIZED_NAME_EVENT_TYPE = "eventType";
  @SerializedName(SERIALIZED_NAME_EVENT_TYPE)
  private V1EventType eventType;

  public static final String SERIALIZED_NAME_MATRIX_KIND = "matrixKind";
  @SerializedName(SERIALIZED_NAME_MATRIX_KIND)
  private V1MatrixKind matrixKind = V1MatrixKind.RANDOM;

  public static final String SERIALIZED_NAME_SCHEDULE_KIND = "scheduleKind";
  @SerializedName(SERIALIZED_NAME_SCHEDULE_KIND)
  private V1ScheduleKind scheduleKind = V1ScheduleKind.CRON;

  public static final String SERIALIZED_NAME_EVENT = "event";
  @SerializedName(SERIALIZED_NAME_EVENT)
  private V1Event event;

  public V1Schemas() {
  }

  public V1Schemas earlyStopping(V1EarlyStopping earlyStopping) {

    this.earlyStopping = earlyStopping;
    return this;
  }

   /**
   * Get earlyStopping
   * @return earlyStopping
  **/
  @javax.annotation.Nullable

  public V1EarlyStopping getEarlyStopping() {
    return earlyStopping;
  }


  public void setEarlyStopping(V1EarlyStopping earlyStopping) {
    this.earlyStopping = earlyStopping;
  }


  public V1Schemas matrix(V1Matrix matrix) {

    this.matrix = matrix;
    return this;
  }

   /**
   * Get matrix
   * @return matrix
  **/
  @javax.annotation.Nullable

  public V1Matrix getMatrix() {
    return matrix;
  }


  public void setMatrix(V1Matrix matrix) {
    this.matrix = matrix;
  }


  public V1Schemas run(V1RunSchema run) {

    this.run = run;
    return this;
  }

   /**
   * Get run
   * @return run
  **/
  @javax.annotation.Nullable

  public V1RunSchema getRun() {
    return run;
  }


  public void setRun(V1RunSchema run) {
    this.run = run;
  }


  public V1Schemas operation(V1Operation operation) {

    this.operation = operation;
    return this;
  }

   /**
   * Get operation
   * @return operation
  **/
  @javax.annotation.Nullable

  public V1Operation getOperation() {
    return operation;
  }


  public void setOperation(V1Operation operation) {
    this.operation = operation;
  }


  public V1Schemas compiledOperation(V1CompiledOperation compiledOperation) {

    this.compiledOperation = compiledOperation;
    return this;
  }

   /**
   * Get compiledOperation
   * @return compiledOperation
  **/
  @javax.annotation.Nullable

  public V1CompiledOperation getCompiledOperation() {
    return compiledOperation;
  }


  public void setCompiledOperation(V1CompiledOperation compiledOperation) {
    this.compiledOperation = compiledOperation;
  }


  public V1Schemas schedule(V1Schedule schedule) {

    this.schedule = schedule;
    return this;
  }

   /**
   * Get schedule
   * @return schedule
  **/
  @javax.annotation.Nullable

  public V1Schedule getSchedule() {
    return schedule;
  }


  public void setSchedule(V1Schedule schedule) {
    this.schedule = schedule;
  }


  public V1Schemas connectionSchema(V1ConnectionSchema connectionSchema) {

    this.connectionSchema = connectionSchema;
    return this;
  }

   /**
   * Get connectionSchema
   * @return connectionSchema
  **/
  @javax.annotation.Nullable

  public V1ConnectionSchema getConnectionSchema() {
    return connectionSchema;
  }


  public void setConnectionSchema(V1ConnectionSchema connectionSchema) {
    this.connectionSchema = connectionSchema;
  }


  public V1Schemas hpParams(V1HpParams hpParams) {

    this.hpParams = hpParams;
    return this;
  }

   /**
   * Get hpParams
   * @return hpParams
  **/
  @javax.annotation.Nullable

  public V1HpParams getHpParams() {
    return hpParams;
  }


  public void setHpParams(V1HpParams hpParams) {
    this.hpParams = hpParams;
  }


  public V1Schemas reference(V1Reference reference) {

    this.reference = reference;
    return this;
  }

   /**
   * Get reference
   * @return reference
  **/
  @javax.annotation.Nullable

  public V1Reference getReference() {
    return reference;
  }


  public void setReference(V1Reference reference) {
    this.reference = reference;
  }


  public V1Schemas artifactsMount(V1ArtifactsMount artifactsMount) {

    this.artifactsMount = artifactsMount;
    return this;
  }

   /**
   * Get artifactsMount
   * @return artifactsMount
  **/
  @javax.annotation.Nullable

  public V1ArtifactsMount getArtifactsMount() {
    return artifactsMount;
  }


  public void setArtifactsMount(V1ArtifactsMount artifactsMount) {
    this.artifactsMount = artifactsMount;
  }


  public V1Schemas polyaxonSidecarContainer(V1PolyaxonSidecarContainer polyaxonSidecarContainer) {

    this.polyaxonSidecarContainer = polyaxonSidecarContainer;
    return this;
  }

   /**
   * Get polyaxonSidecarContainer
   * @return polyaxonSidecarContainer
  **/
  @javax.annotation.Nullable

  public V1PolyaxonSidecarContainer getPolyaxonSidecarContainer() {
    return polyaxonSidecarContainer;
  }


  public void setPolyaxonSidecarContainer(V1PolyaxonSidecarContainer polyaxonSidecarContainer) {
    this.polyaxonSidecarContainer = polyaxonSidecarContainer;
  }


  public V1Schemas polyaxonInitContainer(V1PolyaxonInitContainer polyaxonInitContainer) {

    this.polyaxonInitContainer = polyaxonInitContainer;
    return this;
  }

   /**
   * Get polyaxonInitContainer
   * @return polyaxonInitContainer
  **/
  @javax.annotation.Nullable

  public V1PolyaxonInitContainer getPolyaxonInitContainer() {
    return polyaxonInitContainer;
  }


  public void setPolyaxonInitContainer(V1PolyaxonInitContainer polyaxonInitContainer) {
    this.polyaxonInitContainer = polyaxonInitContainer;
  }


  public V1Schemas artifacs(V1ArtifactsType artifacs) {

    this.artifacs = artifacs;
    return this;
  }

   /**
   * Get artifacs
   * @return artifacs
  **/
  @javax.annotation.Nullable

  public V1ArtifactsType getArtifacs() {
    return artifacs;
  }


  public void setArtifacs(V1ArtifactsType artifacs) {
    this.artifacs = artifacs;
  }


  public V1Schemas wasb(V1WasbType wasb) {

    this.wasb = wasb;
    return this;
  }

   /**
   * Get wasb
   * @return wasb
  **/
  @javax.annotation.Nullable

  public V1WasbType getWasb() {
    return wasb;
  }


  public void setWasb(V1WasbType wasb) {
    this.wasb = wasb;
  }


  public V1Schemas gcs(V1GcsType gcs) {

    this.gcs = gcs;
    return this;
  }

   /**
   * Get gcs
   * @return gcs
  **/
  @javax.annotation.Nullable

  public V1GcsType getGcs() {
    return gcs;
  }


  public void setGcs(V1GcsType gcs) {
    this.gcs = gcs;
  }


  public V1Schemas s3(V1S3Type s3) {

    this.s3 = s3;
    return this;
  }

   /**
   * Get s3
   * @return s3
  **/
  @javax.annotation.Nullable

  public V1S3Type getS3() {
    return s3;
  }


  public void setS3(V1S3Type s3) {
    this.s3 = s3;
  }


  public V1Schemas auth(V1AuthType auth) {

    this.auth = auth;
    return this;
  }

   /**
   * Get auth
   * @return auth
  **/
  @javax.annotation.Nullable

  public V1AuthType getAuth() {
    return auth;
  }


  public void setAuth(V1AuthType auth) {
    this.auth = auth;
  }


  public V1Schemas uri(V1UriType uri) {

    this.uri = uri;
    return this;
  }

   /**
   * Get uri
   * @return uri
  **/
  @javax.annotation.Nullable

  public V1UriType getUri() {
    return uri;
  }


  public void setUri(V1UriType uri) {
    this.uri = uri;
  }


  public V1Schemas resource(V1ConnectionResource resource) {

    this.resource = resource;
    return this;
  }

   /**
   * Get resource
   * @return resource
  **/
  @javax.annotation.Nullable

  public V1ConnectionResource getResource() {
    return resource;
  }


  public void setResource(V1ConnectionResource resource) {
    this.resource = resource;
  }


  public V1Schemas connection(V1ConnectionType connection) {

    this.connection = connection;
    return this;
  }

   /**
   * Get connection
   * @return connection
  **/
  @javax.annotation.Nullable

  public V1ConnectionType getConnection() {
    return connection;
  }


  public void setConnection(V1ConnectionType connection) {
    this.connection = connection;
  }


  public V1Schemas eventType(V1EventType eventType) {

    this.eventType = eventType;
    return this;
  }

   /**
   * Get eventType
   * @return eventType
  **/
  @javax.annotation.Nullable

  public V1EventType getEventType() {
    return eventType;
  }


  public void setEventType(V1EventType eventType) {
    this.eventType = eventType;
  }


  public V1Schemas matrixKind(V1MatrixKind matrixKind) {

    this.matrixKind = matrixKind;
    return this;
  }

   /**
   * Get matrixKind
   * @return matrixKind
  **/
  @javax.annotation.Nullable

  public V1MatrixKind getMatrixKind() {
    return matrixKind;
  }


  public void setMatrixKind(V1MatrixKind matrixKind) {
    this.matrixKind = matrixKind;
  }


  public V1Schemas scheduleKind(V1ScheduleKind scheduleKind) {

    this.scheduleKind = scheduleKind;
    return this;
  }

   /**
   * Get scheduleKind
   * @return scheduleKind
  **/
  @javax.annotation.Nullable

  public V1ScheduleKind getScheduleKind() {
    return scheduleKind;
  }


  public void setScheduleKind(V1ScheduleKind scheduleKind) {
    this.scheduleKind = scheduleKind;
  }


  public V1Schemas event(V1Event event) {

    this.event = event;
    return this;
  }

   /**
   * Get event
   * @return event
  **/
  @javax.annotation.Nullable

  public V1Event getEvent() {
    return event;
  }


  public void setEvent(V1Event event) {
    this.event = event;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1Schemas v1Schemas = (V1Schemas) o;
    return Objects.equals(this.earlyStopping, v1Schemas.earlyStopping) &&
        Objects.equals(this.matrix, v1Schemas.matrix) &&
        Objects.equals(this.run, v1Schemas.run) &&
        Objects.equals(this.operation, v1Schemas.operation) &&
        Objects.equals(this.compiledOperation, v1Schemas.compiledOperation) &&
        Objects.equals(this.schedule, v1Schemas.schedule) &&
        Objects.equals(this.connectionSchema, v1Schemas.connectionSchema) &&
        Objects.equals(this.hpParams, v1Schemas.hpParams) &&
        Objects.equals(this.reference, v1Schemas.reference) &&
        Objects.equals(this.artifactsMount, v1Schemas.artifactsMount) &&
        Objects.equals(this.polyaxonSidecarContainer, v1Schemas.polyaxonSidecarContainer) &&
        Objects.equals(this.polyaxonInitContainer, v1Schemas.polyaxonInitContainer) &&
        Objects.equals(this.artifacs, v1Schemas.artifacs) &&
        Objects.equals(this.wasb, v1Schemas.wasb) &&
        Objects.equals(this.gcs, v1Schemas.gcs) &&
        Objects.equals(this.s3, v1Schemas.s3) &&
        Objects.equals(this.auth, v1Schemas.auth) &&
        Objects.equals(this.uri, v1Schemas.uri) &&
        Objects.equals(this.resource, v1Schemas.resource) &&
        Objects.equals(this.connection, v1Schemas.connection) &&
        Objects.equals(this.eventType, v1Schemas.eventType) &&
        Objects.equals(this.matrixKind, v1Schemas.matrixKind) &&
        Objects.equals(this.scheduleKind, v1Schemas.scheduleKind) &&
        Objects.equals(this.event, v1Schemas.event);
  }

  @Override
  public int hashCode() {
    return Objects.hash(earlyStopping, matrix, run, operation, compiledOperation, schedule, connectionSchema, hpParams, reference, artifactsMount, polyaxonSidecarContainer, polyaxonInitContainer, artifacs, wasb, gcs, s3, auth, uri, resource, connection, eventType, matrixKind, scheduleKind, event);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1Schemas {\n");
    sb.append("    earlyStopping: ").append(toIndentedString(earlyStopping)).append("\n");
    sb.append("    matrix: ").append(toIndentedString(matrix)).append("\n");
    sb.append("    run: ").append(toIndentedString(run)).append("\n");
    sb.append("    operation: ").append(toIndentedString(operation)).append("\n");
    sb.append("    compiledOperation: ").append(toIndentedString(compiledOperation)).append("\n");
    sb.append("    schedule: ").append(toIndentedString(schedule)).append("\n");
    sb.append("    connectionSchema: ").append(toIndentedString(connectionSchema)).append("\n");
    sb.append("    hpParams: ").append(toIndentedString(hpParams)).append("\n");
    sb.append("    reference: ").append(toIndentedString(reference)).append("\n");
    sb.append("    artifactsMount: ").append(toIndentedString(artifactsMount)).append("\n");
    sb.append("    polyaxonSidecarContainer: ").append(toIndentedString(polyaxonSidecarContainer)).append("\n");
    sb.append("    polyaxonInitContainer: ").append(toIndentedString(polyaxonInitContainer)).append("\n");
    sb.append("    artifacs: ").append(toIndentedString(artifacs)).append("\n");
    sb.append("    wasb: ").append(toIndentedString(wasb)).append("\n");
    sb.append("    gcs: ").append(toIndentedString(gcs)).append("\n");
    sb.append("    s3: ").append(toIndentedString(s3)).append("\n");
    sb.append("    auth: ").append(toIndentedString(auth)).append("\n");
    sb.append("    uri: ").append(toIndentedString(uri)).append("\n");
    sb.append("    resource: ").append(toIndentedString(resource)).append("\n");
    sb.append("    connection: ").append(toIndentedString(connection)).append("\n");
    sb.append("    eventType: ").append(toIndentedString(eventType)).append("\n");
    sb.append("    matrixKind: ").append(toIndentedString(matrixKind)).append("\n");
    sb.append("    scheduleKind: ").append(toIndentedString(scheduleKind)).append("\n");
    sb.append("    event: ").append(toIndentedString(event)).append("\n");
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
    openapiFields.add("earlyStopping");
    openapiFields.add("matrix");
    openapiFields.add("run");
    openapiFields.add("operation");
    openapiFields.add("compiledOperation");
    openapiFields.add("schedule");
    openapiFields.add("connectionSchema");
    openapiFields.add("hpParams");
    openapiFields.add("reference");
    openapiFields.add("artifactsMount");
    openapiFields.add("polyaxonSidecarContainer");
    openapiFields.add("polyaxonInitContainer");
    openapiFields.add("artifacs");
    openapiFields.add("wasb");
    openapiFields.add("gcs");
    openapiFields.add("s3");
    openapiFields.add("auth");
    openapiFields.add("uri");
    openapiFields.add("resource");
    openapiFields.add("connection");
    openapiFields.add("eventType");
    openapiFields.add("matrixKind");
    openapiFields.add("scheduleKind");
    openapiFields.add("event");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to V1Schemas
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!V1Schemas.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in V1Schemas is not found in the empty JSON string", V1Schemas.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!V1Schemas.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `V1Schemas` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }
      // validate the optional field `earlyStopping`
      if (jsonObj.get("earlyStopping") != null && !jsonObj.get("earlyStopping").isJsonNull()) {
        V1EarlyStopping.validateJsonObject(jsonObj.getAsJsonObject("earlyStopping"));
      }
      // validate the optional field `matrix`
      if (jsonObj.get("matrix") != null && !jsonObj.get("matrix").isJsonNull()) {
        V1Matrix.validateJsonObject(jsonObj.getAsJsonObject("matrix"));
      }
      // validate the optional field `run`
      if (jsonObj.get("run") != null && !jsonObj.get("run").isJsonNull()) {
        V1RunSchema.validateJsonObject(jsonObj.getAsJsonObject("run"));
      }
      // validate the optional field `operation`
      if (jsonObj.get("operation") != null && !jsonObj.get("operation").isJsonNull()) {
        V1Operation.validateJsonObject(jsonObj.getAsJsonObject("operation"));
      }
      // validate the optional field `compiledOperation`
      if (jsonObj.get("compiledOperation") != null && !jsonObj.get("compiledOperation").isJsonNull()) {
        V1CompiledOperation.validateJsonObject(jsonObj.getAsJsonObject("compiledOperation"));
      }
      // validate the optional field `schedule`
      if (jsonObj.get("schedule") != null && !jsonObj.get("schedule").isJsonNull()) {
        V1Schedule.validateJsonObject(jsonObj.getAsJsonObject("schedule"));
      }
      // validate the optional field `connectionSchema`
      if (jsonObj.get("connectionSchema") != null && !jsonObj.get("connectionSchema").isJsonNull()) {
        V1ConnectionSchema.validateJsonObject(jsonObj.getAsJsonObject("connectionSchema"));
      }
      // validate the optional field `hpParams`
      if (jsonObj.get("hpParams") != null && !jsonObj.get("hpParams").isJsonNull()) {
        V1HpParams.validateJsonObject(jsonObj.getAsJsonObject("hpParams"));
      }
      // validate the optional field `reference`
      if (jsonObj.get("reference") != null && !jsonObj.get("reference").isJsonNull()) {
        V1Reference.validateJsonObject(jsonObj.getAsJsonObject("reference"));
      }
      // validate the optional field `artifactsMount`
      if (jsonObj.get("artifactsMount") != null && !jsonObj.get("artifactsMount").isJsonNull()) {
        V1ArtifactsMount.validateJsonObject(jsonObj.getAsJsonObject("artifactsMount"));
      }
      // validate the optional field `polyaxonSidecarContainer`
      if (jsonObj.get("polyaxonSidecarContainer") != null && !jsonObj.get("polyaxonSidecarContainer").isJsonNull()) {
        V1PolyaxonSidecarContainer.validateJsonObject(jsonObj.getAsJsonObject("polyaxonSidecarContainer"));
      }
      // validate the optional field `polyaxonInitContainer`
      if (jsonObj.get("polyaxonInitContainer") != null && !jsonObj.get("polyaxonInitContainer").isJsonNull()) {
        V1PolyaxonInitContainer.validateJsonObject(jsonObj.getAsJsonObject("polyaxonInitContainer"));
      }
      // validate the optional field `artifacs`
      if (jsonObj.get("artifacs") != null && !jsonObj.get("artifacs").isJsonNull()) {
        V1ArtifactsType.validateJsonObject(jsonObj.getAsJsonObject("artifacs"));
      }
      // validate the optional field `wasb`
      if (jsonObj.get("wasb") != null && !jsonObj.get("wasb").isJsonNull()) {
        V1WasbType.validateJsonObject(jsonObj.getAsJsonObject("wasb"));
      }
      // validate the optional field `gcs`
      if (jsonObj.get("gcs") != null && !jsonObj.get("gcs").isJsonNull()) {
        V1GcsType.validateJsonObject(jsonObj.getAsJsonObject("gcs"));
      }
      // validate the optional field `s3`
      if (jsonObj.get("s3") != null && !jsonObj.get("s3").isJsonNull()) {
        V1S3Type.validateJsonObject(jsonObj.getAsJsonObject("s3"));
      }
      // validate the optional field `auth`
      if (jsonObj.get("auth") != null && !jsonObj.get("auth").isJsonNull()) {
        V1AuthType.validateJsonObject(jsonObj.getAsJsonObject("auth"));
      }
      // validate the optional field `uri`
      if (jsonObj.get("uri") != null && !jsonObj.get("uri").isJsonNull()) {
        V1UriType.validateJsonObject(jsonObj.getAsJsonObject("uri"));
      }
      // validate the optional field `resource`
      if (jsonObj.get("resource") != null && !jsonObj.get("resource").isJsonNull()) {
        V1ConnectionResource.validateJsonObject(jsonObj.getAsJsonObject("resource"));
      }
      // validate the optional field `connection`
      if (jsonObj.get("connection") != null && !jsonObj.get("connection").isJsonNull()) {
        V1ConnectionType.validateJsonObject(jsonObj.getAsJsonObject("connection"));
      }
      // validate the optional field `eventType`
      if (jsonObj.get("eventType") != null && !jsonObj.get("eventType").isJsonNull()) {
        V1EventType.validateJsonObject(jsonObj.getAsJsonObject("eventType"));
      }
      // validate the optional field `event`
      if (jsonObj.get("event") != null && !jsonObj.get("event").isJsonNull()) {
        V1Event.validateJsonObject(jsonObj.getAsJsonObject("event"));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!V1Schemas.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'V1Schemas' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<V1Schemas> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(V1Schemas.class));

       return (TypeAdapter<T>) new TypeAdapter<V1Schemas>() {
           @Override
           public void write(JsonWriter out, V1Schemas value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public V1Schemas read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of V1Schemas given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of V1Schemas
  * @throws IOException if the JSON string is invalid with respect to V1Schemas
  */
  public static V1Schemas fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, V1Schemas.class);
  }

 /**
  * Convert an instance of V1Schemas to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

