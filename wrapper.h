#include "nitro.h"
#include "frame.h"
#include "err.h"
int nitro_send_(nitro_frame_t *fr, nitro_socket_t *s, int flags)
{
  int out = nitro_send(&fr, s, flags);
  return (out);
}

nitro_frame_t * nitro_recv_(nitro_socket_t *s, int flags)
{
  nitro_frame_t * out = nitro_recv(s, flags);
  return (out);
}

int nitro_reply_(nitro_frame_t *snd, nitro_frame_t *fr, nitro_socket_t *s, int flags)
{
  int out = nitro_reply(snd, &fr, s, flags);
  return (out);
}

int nitro_relay_fw_(nitro_frame_t *snd, nitro_frame_t *fr, nitro_socket_t *s, int flags)
{
  int out = nitro_relay_fw(snd, &fr, s, flags);
  return (out);
}

int nitro_relay_bk_(nitro_frame_t *snd, nitro_frame_t *fr, nitro_socket_t *s, int flags)
{
  int out = nitro_relay_bk(snd, &fr, s, flags);
  return (out);
}

int nitro_sub_(nitro_socket_t *s, uint8_t *key, size_t length)
{
  int out = nitro_sub(s, key, length);
  return (out);
}

int nitro_unsub_(nitro_socket_t *s, uint8_t *key, size_t length)
{
  int out = nitro_unsub(s, key, length);
  return (out);
}

int nitro_pub_(nitro_frame_t *fr, uint8_t *key, size_t length, nitro_socket_t * s, int flags)
{
  int out = nitro_pub(&fr, key, length, s, flags);
  return (out);
}

int nitro_eventfd_(nitro_socket_t *s)
{
  return (nitro_eventfd(s));
}

void nitro_frame_destroy_(nitro_frame_t *f)
{
  return (nitro_frame_destroy(f));
}


