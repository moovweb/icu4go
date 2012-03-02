#include "helper.h"
#include <string.h>
#include <unicode/utypes.h>
#include <unicode/ucsdet.h>

char* detectCharset(void *detector, void *input, int input_len, int *confidence, int *status) {
	const UCharsetMatch **matches;
	int matchCount, i, conf;
	const char *charsetName;
	const char *bestGuessedCharset;

	
	ucsdet_setText((UCharsetDetector*)detector, (char*)input, input_len, status);
	if (*status != U_ZERO_ERROR) {
		return NULL;
	}
	matches = ucsdet_detectAll((UCharsetDetector*)detector, &matchCount, status);
	if (*status != U_ZERO_ERROR) {
		return NULL;
	}
	if (matchCount > 0) {
		bestGuessedCharset = ucsdet_getName(matches[0], status);
		if (*status != U_ZERO_ERROR) {
			return NULL;
		}
		conf = ucsdet_getConfidence(matches[0], status);
		if (*status != U_ZERO_ERROR) {
			return NULL;
		}
		*confidence = conf;
	}
	return bestGuessedCharset;
}