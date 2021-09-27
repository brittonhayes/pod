import {
  Limiter,
  Meter,
  Player,
  ToneAudioBuffer,
  Distortion,
  Reverb,
  Filter,
  FilterRollOff,
} from "tone";
import { Frequency, Seconds } from "tone/build/esm/core/type/Units";

export const MIN_FREQ = 20;
export const MAX_FREQ = 20000;

/**
 * Processor wraps a Player with a set of available
 * DSP methods
 */
export class Processor {
  player: Player;
  debug?: boolean;
  meter?: Meter;

  constructor(player: Player, debug?: boolean) {
    if (debug) {
      this.debug = debug;
      player.debug = debug;
    }

    this.player = player;
  }

  public Play(blob: Blob): void {
    const audioContext = new AudioContext();
    const fileReader = new FileReader();

    if (!blob.type.includes("audio")) {
      console.error("cannot create audio plaer for non-audio blob");
      return;
    }

    if (this.debug) {
      fileReader.onloadstart = () => {
        console.debug("started loading audio buffer");
      };
    }

    fileReader.onloadend = () => {
      if (this.debug) {
        console.info("loading done");
      }
      const arrayBuffer = fileReader.result as ArrayBuffer;
      audioContext.decodeAudioData(
        arrayBuffer,
        (audioBuffer) => {
          this.player.buffer = new ToneAudioBuffer(audioBuffer);
          this.player.toDestination().stop();
          this.player.toDestination().start();
        },
        (e) => {
          console.error(e);
          return;
        }
      );
    };
    fileReader.readAsArrayBuffer(blob);
  }

  public WithFilter(
    freq?: Frequency,
    type?: BiquadFilterType,
    rolloff?: FilterRollOff
  ): Processor {
    const filter = new Filter(freq, type, rolloff).toDestination();
    this.player.connect(filter);

    return this;
  }

  public WithReverb(decay: Seconds): Processor {
    const reverb = new Reverb(decay).toDestination();
    this.player.connect(reverb);

    return this;
  }

  /**
   * WithLimiter returns a processor with
   * a level limiter to control peaks
   * @param peak
   * @returns Processor
   */
  public WithLimiter(peak: number): Processor {
    const defaultPeak = -1;
    if (peak < -100 || peak > 0) {
      console.warn(`RangeError: Value must be within [-100, 0], got: ${peak}`);
      peak = defaultPeak;
    }

    const limiter = new Limiter(peak);
    this.player.connect(limiter).toDestination();

    return this;
  }

  public WithDistortion(intensity: number): Processor {
    if (intensity < -100 || intensity > 0) {
      console.warn(
        `RangeError: Value must be within [0, 1], got: ${intensity}`
      );
      intensity = 0.5;
    }

    const dist = new Distortion(intensity).toDestination();
    this.player.connect(dist);

    return this;
  }

  /**
   * WithMeter returns a processor with
   * an RMS meter
   * @returns Processor
   */
  public WithMeter(): Processor {
    const meter = new Meter();
    this.meter = meter;
    this.player.connect(meter).toDestination();

    return this;
  }
}
