package deepSpeach

/*
#cgo CFLAGS: -std=c99
#cgo CXXFLAGS: -I/usr/lib/
#cgo LDFLAGS: -L/usr/lib/ -lstdc++

#include "brigde/deepspeech.h"
#include "brigde/args.h"

void
ProcessFile(ModelState* context, const char* path, bool show_times)
{
	ds_audio_buffer audio = GetAudioBuffer(path, DS_GetModelSampleRate(context));

	// Pass audio to DeepSpeech
	// We take half of buffer_size because buffer is a char* while
	// LocalDsSTT() expected a short*
	ds_result result = LocalDsSTT(context,
							(const short*)audio.buffer,
							audio.buffer_size / 2,
							extended_metadata,
							json_output);
	free(audio.buffer);

	if (result.string) {
		printf("%s\n", result.string);
		DS_FreeString((char*)result.string);
	}

	if (show_times) {
		printf("cpu_time_overall=%.05f\n",
		result.cpu_time_overall);
	}
}

typedef struct {
  char*  buffer;
  size_t buffer_size;
} ds_audio_buffer;

ds_audio_buffer
GetAudioBuffer(const char* path, int desired_sample_rate)
{
  ds_audio_buffer res = {0};

  // FIXME: Hack and support only mono 16-bits PCM with standard SoX header
  FILE* wave = fopen(path, "r");

  size_t rv;

  unsigned short audio_format;
  fseek(wave, 20, SEEK_SET); rv = fread(&audio_format, 2, 1, wave);

  unsigned short num_channels;
  fseek(wave, 22, SEEK_SET); rv = fread(&num_channels, 2, 1, wave);

  unsigned int sample_rate;
  fseek(wave, 24, SEEK_SET); rv = fread(&sample_rate, 4, 1, wave);

  unsigned short bits_per_sample;
  fseek(wave, 34, SEEK_SET); rv = fread(&bits_per_sample, 2, 1, wave);

  assert(audio_format == 1); // 1 is PCM
  assert(num_channels == 1); // MONO
  assert(sample_rate == desired_sample_rate); // at desired sample rate
  assert(bits_per_sample == 16); // 16 bits per sample

  fprintf(stderr, "audio_format=%d\n", audio_format);
  fprintf(stderr, "num_channels=%d\n", num_channels);
  fprintf(stderr, "sample_rate=%d (desired=%d)\n", sample_rate, desired_sample_rate);
  fprintf(stderr, "bits_per_sample=%d\n", bits_per_sample);

  fseek(wave, 40, SEEK_SET); rv = fread(&res.buffer_size, 4, 1, wave);
  fprintf(stderr, "res.buffer_size=%ld\n", res.buffer_size);

  fseek(wave, 44, SEEK_SET);
  res.buffer = (char*)malloc(sizeof(char) * res.buffer_size);
  rv = fread(res.buffer, sizeof(char), res.buffer_size, wave);

  fclose(wave);

  return res;
}

ds_result
LocalDsSTT(ModelState* aCtx, const short* aBuffer, size_t aBufferSize,
           bool extended_output, bool json_output)
{
  ds_result res = {0};

  clock_t ds_start_time = clock();

  // sphinx-doc: c_ref_inference_start
  if (extended_output) {
    Metadata *result = DS_SpeechToTextWithMetadata(aCtx, aBuffer, aBufferSize, 1);
    res.string = CandidateTranscriptToString(&result->transcripts[0]);
    DS_FreeMetadata(result);
  } else if (json_output) {
    Metadata *result = DS_SpeechToTextWithMetadata(aCtx, aBuffer, aBufferSize, json_candidate_transcripts);
    res.string = MetadataToJSON(result);
    DS_FreeMetadata(result);
  } else if (stream_size > 0) {
    StreamingState* ctx;
    int status = DS_CreateStream(aCtx, &ctx);
    if (status != DS_ERR_OK) {
      res.string = strdup("");
      return res;
    }
    size_t off = 0;
    const char *last = nullptr;
    const char *prev = nullptr;
    while (off < aBufferSize) {
      size_t cur = aBufferSize - off > stream_size ? stream_size : aBufferSize - off;
      DS_FeedAudioContent(ctx, aBuffer + off, cur);
      off += cur;
      prev = last;
      const char* partial = DS_IntermediateDecode(ctx);
      if (last == nullptr || strcmp(last, partial)) {
        printf("%s\n", partial);
        last = partial;
      } else {
        DS_FreeString((char *) partial);
      }
      if (prev != nullptr && prev != last) {
        DS_FreeString((char *) prev);
      }
    }
    if (last != nullptr) {
      DS_FreeString((char *) last);
    }
    res.string = DS_FinishStream(ctx);
  } else if (extended_stream_size > 0) {
    StreamingState* ctx;
    int status = DS_CreateStream(aCtx, &ctx);
    if (status != DS_ERR_OK) {
      res.string = strdup("");
      return res;
    }
    size_t off = 0;
    const char *last = nullptr;
    const char *prev = nullptr;
    while (off < aBufferSize) {
      size_t cur = aBufferSize - off > extended_stream_size ? extended_stream_size : aBufferSize - off;
      DS_FeedAudioContent(ctx, aBuffer + off, cur);
      off += cur;
      prev = last;
      const Metadata* result = DS_IntermediateDecodeWithMetadata(ctx, 1);
      const char* partial = CandidateTranscriptToString(&result->transcripts[0]);
      if (last == nullptr || strcmp(last, partial)) {
        printf("%s\n", partial);
       last = partial;
      } else {
        free((char *) partial);
      }
      if (prev != nullptr && prev != last) {
        free((char *) prev);
      }
      DS_FreeMetadata((Metadata *)result);
    }
    const Metadata* result = DS_FinishStreamWithMetadata(ctx, 1);
    res.string = CandidateTranscriptToString(&result->transcripts[0]);
    DS_FreeMetadata((Metadata *)result);
    free((char *) last);
  } else {
    res.string = DS_SpeechToText(aCtx, aBuffer, aBufferSize);
  }
  // sphinx-doc: c_ref_inference_stop

  clock_t ds_end_infer = clock();

  res.cpu_time_overall =
    ((double) (ds_end_infer - ds_start_time)) / CLOCKS_PER_SEC;

  return res;
}
*/
import "C"
import (
	"flag"
	"fmt"
)

func main() {
	model := flag.String("model", "", "Path to the model (protocol buffer binary file)")
	scorer := flag.String("scorer", "", "Path to the external scorer file")
	audio := flag.String("audio", "", "Path to the audio file to run (WAV format)")
	flag.Parse()

	var ctx *C.ModelState
	var status int

	status = C.DS_CreateModel(model, &ctx)
	if status != 0 {
		panic(fmt.Errorf("could not create model"))
	}

	status = C.DS_EnableExternalScorer(ctx, scorer)
	if status != 0 {
		panic(fmt.Errorf("could not enable external scorer"))
	}

	C.ProcessFile(ctx, audio, false)

	C.DS_FreeModel(ctx)
}
