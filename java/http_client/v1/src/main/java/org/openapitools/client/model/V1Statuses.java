/*
 * Polyaxon SDKs and REST API specification.
 *    
 *
 * The version of the OpenAPI document: 2.0.0-rc16
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.annotations.SerializedName;

import java.io.IOException;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;

/**
 * Gets or Sets v1Statuses
 */
@JsonAdapter(V1Statuses.Adapter.class)
public enum V1Statuses {
  
  CREATED("created"),
  
  RESUMING("resuming"),
  
  ON_SCHEDULE("on_schedule"),
  
  COMPILED("compiled"),
  
  QUEUED("queued"),
  
  SCHEDULED("scheduled"),
  
  STARTING("starting"),
  
  RUNNING("running"),
  
  PROCESSING("processing"),
  
  STOPPING("stopping"),
  
  FAILED("failed"),
  
  STOPPED("stopped"),
  
  SUCCEEDED("succeeded"),
  
  SKIPPED("skipped"),
  
  WARNING("warning"),
  
  UNSCHEDULABLE("unschedulable"),
  
  UPSTREAM_FAILED("upstream_failed"),
  
  RETRYING("retrying"),
  
  UNKNOWN("unknown"),
  
  DONE("done");

  private String value;

  V1Statuses(String value) {
    this.value = value;
  }

  public String getValue() {
    return value;
  }

  @Override
  public String toString() {
    return String.valueOf(value);
  }

  public static V1Statuses fromValue(String value) {
    for (V1Statuses b : V1Statuses.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }

  public static class Adapter extends TypeAdapter<V1Statuses> {
    @Override
    public void write(final JsonWriter jsonWriter, final V1Statuses enumeration) throws IOException {
      jsonWriter.value(enumeration.getValue());
    }

    @Override
    public V1Statuses read(final JsonReader jsonReader) throws IOException {
      String value = jsonReader.nextString();
      return V1Statuses.fromValue(value);
    }
  }
}

