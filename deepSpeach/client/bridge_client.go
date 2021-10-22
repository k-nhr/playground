package client

/*
#cgo CFLAGS: -x c++ -fpermissive -I/usr/include/c++/9/ -I/usr/include/x86_64-linux-gnu/c++/9 -I./ctcdecode/third_party/object_pool -I./ctcdecode/third_party/ThreadPool -I./kenlm
#cgo LDFLAGS: -L/usr/lib/ -lrt -lfst -lstdc++

#include <assert.h>
#include <time.h>

#include "deepspeech.h"
#include "args.h"


typedef struct {
  const char* string;
  double cpu_time_overall;
} ds_result;

ds_result
LocalDsSTT(ModelState* aCtx, const short* aBuffer, size_t aBufferSize)
{
  ds_result res = {0};

  clock_t ds_start_time = clock();
  res.string = DS_SpeechToText(aCtx, aBuffer, aBufferSize);
  clock_t ds_end_infer = clock();

  res.cpu_time_overall =
    ((double) (ds_end_infer - ds_start_time)) / CLOCKS_PER_SEC;

  return res;
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
void
ProcessFile(ModelState* context, const char* path)
{
  ds_audio_buffer audio = GetAudioBuffer(path, DS_GetModelSampleRate(context));

  // Pass audio to DeepSpeech
  // We take half of buffer_size because buffer is a char* while
  // LocalDsSTT() expected a short*
  ds_result result = LocalDsSTT(context, (const short*)audio.buffer, audio.buffer_size / 2);
  free(audio.buffer);

  if (result.string) {
    printf("%s\n", result.string);
    DS_FreeString((char*)result.string);
  }
}
*/
import "C"
import (
	"fmt"
)

var ctx *C.ModelState

func CreateModel(model string) error {
	sts := C.DS_CreateModel(C.CString(model), &ctx)
	if sts != 0 {
		s := C.GoString(C.DS_ErrorCodeToErrorMessage(sts))
		return fmt.Errorf("could not create model: %s", s)
	}
	return nil
}

func EnableExternalScorer(scorer string) error {
	sts := C.DS_EnableExternalScorer(ctx, C.CString(scorer))
	if sts != 0 {
		return fmt.Errorf("could not enable external scorer")
	}
	return nil
}

func ProcessFile(audio string) {
	C.ProcessFile(ctx, C.CString(audio))
	return
}

func FreeModel() {
	C.DS_FreeModel(ctx)
	return
}
